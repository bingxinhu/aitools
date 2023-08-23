package apis

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go-admin/common/apis"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Test1 struct {
	apis.Api
}

func (e Test1) GetTest1(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}

	e.OK("Hello test1!", "success")
}
func getFileMd5(filePath string) string {
	pFile, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("打开文件失败：%s", err)
		return ""
	}
	defer pFile.Close()
	md5h := md5.New()
	io.Copy(md5h, pFile)

	return hex.EncodeToString(md5h.Sum(nil))
}

func (e Test1) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "读取失败: "+err.Error())
		return
	}
	var uploadir string
	//定义文件保存地址
	uploadir = "./static/myuploadfiles/"
	_, err = os.Stat(uploadir)
	if os.IsNotExist(err) {
		os.Mkdir(uploadir, os.ModePerm)
	}
	//fileName 脱敏
	fileId := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000)
	newFileName := fileId + path.Ext(file.Filename)
	dst := uploadir + newFileName
	uplouderr := c.SaveUploadedFile(file, dst)
	if uplouderr != nil {
		e.OK("null", "success")
		//c.JSON(http.StatusOK, gin.H{"filename": "nil"})
		return
	}
	mdtstr := getFileMd5(dst) + path.Ext(file.Filename)
	os.Rename(dst, uploadir+mdtstr)
	err = e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(mdtstr, "success")
	//c.JSON(http.StatusOK, gin.H{"filename": newFileName})
}

func (e Test1) UploadFile_definename(c *gin.Context) {
	newFileName := c.PostForm("savename")
	if newFileName == "" {
		c.String(http.StatusInternalServerError, "savename get failed!")
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "读取失败: "+err.Error())
		return
	}
	var uploadir string
	//定义文件保存地址
	uploadir = "./static/myuploadfiles/"
	_, err = os.Stat(uploadir)
	if os.IsNotExist(err) {
		os.Mkdir(uploadir, os.ModePerm)
	}
	dst := uploadir + newFileName
	uplouderr := c.SaveUploadedFile(file, dst)
	if uplouderr != nil {
		e.OK("null", "success")
		return
	}
	err = e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(newFileName, "success")
}
