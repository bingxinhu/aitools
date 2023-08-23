#ifndef CURLFUNC_HPP
#define CURLFUNC_HPP
#include<iostream>
#include<string>
#include<curl/curl.h>
using namespace std;
//get请求和post请求数据响应函数
size_t req_reply_post(void *ptr, size_t size, size_t nmemb, void *stream)
{
    //在注释的里面可以打印请求流，cookie的信息
    cout << "----->reply" << endl;
    string *str = (string*)stream;
    //cout << *str << endl;
    (*str).append((char*)ptr, size*nmemb);
    return size * nmemb;
}
//http POST请求
CURLcode curl_post_req(const string &url, const string &postParams, string &response)
{
    // curl初始化
    CURL *curl = curl_easy_init();
    // curl返回值
    CURLcode res;
    if (curl)
    {
        // set params
        //设置curl的请求头
        struct curl_slist* header_list = NULL;
        header_list = curl_slist_append(header_list, "User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko");
        header_list = curl_slist_append(header_list, "Content-Type:application/x-www-form-urlencoded; charset=UTF-8");
        curl_easy_setopt(curl, CURLOPT_HTTPHEADER, header_list);

        //不接收响应头数据0代表不接收 1代表接收
        curl_easy_setopt(curl, CURLOPT_HEADER, 0);

        //设置请求为post请求
        curl_easy_setopt(curl, CURLOPT_POST, 1);

        //设置请求的URL地址
        curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
        //设置post请求的参数
        curl_easy_setopt(curl, CURLOPT_POSTFIELDS, postParams.c_str());

        //设置ssl验证
        curl_easy_setopt(curl, CURLOPT_SSL_VERIFYPEER, false);
        curl_easy_setopt(curl, CURLOPT_SSL_VERIFYHOST, false);

        //CURLOPT_VERBOSE的值为1时，会显示详细的调试信息
        curl_easy_setopt(curl, CURLOPT_VERBOSE, 0);

        curl_easy_setopt(curl, CURLOPT_READFUNCTION, NULL);

        //设置数据接收和写入函数
        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, req_reply_post);
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, (void *)&response);

        curl_easy_setopt(curl, CURLOPT_NOSIGNAL, 1);

        //设置超时时间
        curl_easy_setopt(curl, CURLOPT_CONNECTTIMEOUT, 6);
        curl_easy_setopt(curl, CURLOPT_TIMEOUT, 6);

        // 开启post请求
        res = curl_easy_perform(curl);
    }
    //释放curl
    curl_easy_cleanup(curl);
    return res;
}

size_t req_reply_get(void *ptr, size_t size, size_t nmemb, void *stream)
{
    //在注释的里面可以打印请求流，cookie的信息
    //cout << "----->reply" << endl;
    string *str = (string*)stream;
    //cout << *str << endl;
    (*str).append((char*)ptr, size*nmemb);
    return size * nmemb;
}
//http GET请求
CURLcode curl_get_req(const std::string &url, std::string &response)
{
    //curl初始化
    CURL *curl = curl_easy_init();
    // curl返回值
    CURLcode res;
    if (curl)
    {
        //设置curl的请求头
        struct curl_slist* header_list = NULL;
        header_list = curl_slist_append(header_list, "User-Agent: Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko");
        curl_easy_setopt(curl, CURLOPT_HTTPHEADER, header_list);

        //不接收响应头数据0代表不接收 1代表接收
        curl_easy_setopt(curl, CURLOPT_HEADER, 0);

        //设置请求的URL地址
        curl_easy_setopt(curl, CURLOPT_URL, url.c_str());

        //设置ssl验证
        curl_easy_setopt(curl, CURLOPT_SSL_VERIFYPEER, false);
        curl_easy_setopt(curl, CURLOPT_SSL_VERIFYHOST, false);

        //CURLOPT_VERBOSE的值为1时，会显示详细的调试信息
        curl_easy_setopt(curl, CURLOPT_VERBOSE, 0);

        curl_easy_setopt(curl, CURLOPT_READFUNCTION, NULL);

        //设置数据接收函数
        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, req_reply_get);
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, (void *)&response);

        curl_easy_setopt(curl, CURLOPT_NOSIGNAL, 1);

        //设置超时时间
        curl_easy_setopt(curl, CURLOPT_CONNECTTIMEOUT, 6); // set transport and time out time
        curl_easy_setopt(curl, CURLOPT_TIMEOUT, 6);

        // 开启请求
        res = curl_easy_perform(curl);
    }
    // 释放curl
    curl_easy_cleanup(curl);
    return res;
}

std::string curl_shell_put(std::string cmd) {
    FILE *fp = NULL;
    char buf[1024];
    std::string rstr;
    if( (fp = popen(cmd.c_str(), "r")) != NULL)
    {
        while(fgets(buf, 1024, fp) != NULL)
        {
            rstr += std::string(buf);
        }
        pclose(fp);
        fp = NULL;
    }
    return rstr;
}

size_t writedata2file(void *ptr, size_t size, size_t nmemb, FILE *stream) {
    size_t written = fwrite(ptr, size, nmemb, stream);
    return written;
}

int curl_download_file(const char* url, const char* outfilename){
    CURL *curl;
    FILE *fp;
    CURLcode res;
    /*   调用curl_global_init()初始化libcurl  */
    res = curl_global_init(CURL_GLOBAL_ALL);
    if (CURLE_OK != res)
    {
        printf("init libcurl failed.");
        curl_global_cleanup();
        return -1;
    }
    /*  调用curl_easy_init()函数得到 easy interface型指针  */
    curl = curl_easy_init();
    if (curl) {
        fp = fopen(outfilename,"wb");
 
        /*  调用curl_easy_setopt()设置传输选项 */
        res = curl_easy_setopt(curl, CURLOPT_URL, url);
        if (res != CURLE_OK)
        {
            fclose(fp);
            curl_easy_cleanup(curl);
            return -1;
        }
        /*  根据curl_easy_setopt()设置的传输选项，实现回调函数以完成用户特定任务  */
        res = curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, writedata2file);
        if (res != CURLE_OK){
            fclose(fp);
            curl_easy_cleanup(curl);
            return -1;
        }
        /*  根据curl_easy_setopt()设置的传输选项，实现回调函数以完成用户特定任务  */
        res = curl_easy_setopt(curl, CURLOPT_WRITEDATA, fp);
        if (res != CURLE_OK)
        {
            fclose(fp);
            curl_easy_cleanup(curl);
            return -1;
        }
 
        res = curl_easy_perform(curl);
        // 调用curl_easy_perform()函数完成传输任务
        fclose(fp);
        /* Check for errors */
        if(res != CURLE_OK){
            fprintf(stderr, "curl_easy_perform() failed: %s\n",curl_easy_strerror(res));
            curl_easy_cleanup(curl);
            return -1;
        }
 
        /* always cleanup */
        curl_easy_cleanup(curl);
        // 调用curl_easy_cleanup()释放内存
 
    }
    curl_global_cleanup();
    return 0;
}

int write_data(void* buffer, int size, int nmemb, void* userp){
    std::string * str = dynamic_cast<std::string *>((std::string *)userp);
    str->append((char *)buffer, size * nmemb);
    return nmemb;
}
int curl_upload_file(string url, std::string srcpath, std::string dstpath, string &body,  string* response)
{
    CURL *curl;
    CURLcode ret;
    curl = curl_easy_init();
    struct curl_httppost* post = NULL;
    struct curl_httppost* last = NULL;
    if (curl)
    {
        curl_easy_setopt(curl, CURLOPT_URL, (char *)url.c_str());           //指定url
        curl_formadd(&post, &last, CURLFORM_PTRNAME, "savename", CURLFORM_PTRCONTENTS, /*"device_cover"*/dstpath.c_str(), CURLFORM_END);//form-data key(path) 和 value(device_cover)
        curl_formadd(&post, &last, CURLFORM_PTRNAME,  "file", CURLFORM_FILE, srcpath.c_str(), CURLFORM_FILENAME, dstpath.c_str(), CURLFORM_END);// form-data key(file) "./test.jpg"为文件路径  "hello.jpg" 为文件上传时文件名
        curl_easy_setopt(curl, CURLOPT_HTTPPOST, post);                     //构造post参数    
        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_data);          //绑定相应
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, (void *)response);        //绑定响应内容的地址

        ret = curl_easy_perform(curl);                          //执行请求
        if(ret == 0){
            curl_easy_cleanup(curl);    
            return 0;  
        }
        else{
            return ret;
        }
    }
	else{
        return -1;
	}
}

#endif // CURLFUNC_HPP
