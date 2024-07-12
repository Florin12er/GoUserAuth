package controllers

import (
    "net/http"
    "os"
    "time"
    "userAuth/initializers"
    "userAuth/models" // Assuming you have a models package with a User struct

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
    var body struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.BindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "failed to read body",
        })
        return
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "failed to hash password",
        })
        return
    }

    user := models.User{
        Name:     body.Name,
        Email:    body.Email,
        Password: string(hash),
    }

    if err := initializers.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "failed to create user",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "user created successfully",
    })
}

func Login(c *gin.Context) {
    var body struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.BindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "failed to read body",
        })
        return
    }

    var user models.User
    if err := initializers.DB.First(&user, "email = ?", body.Email).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "invalid email or password",
        })
        return
    }

    if user.ID == 0 {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "invalid email or password",
        })
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "invalid email or password",
        })
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": user.ID,
        "exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
    })
    tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "failed to create token",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token": tokenString,
    })
}

