#include "TaskDealer.h"
#include "curlfunc.hpp"

TaskDealer::TaskDealer(PlatformType platform) : _platform(platform)
{

}

bool TaskDealer::Init(const char* hist_ip) {
    _host_ip = hist_ip;
    _query_url = std::string("http://") + _host_ip + std::string(":8000/api/v1/model-compile-apply/query_noauth?status=待编译");
    _update_url = std::string("http://") + _host_ip + std::string(":8000/api/v1/model-compile-apply/update_noauth/");
    _download_url = std::string("http://") + _host_ip +std::string(":8000/static/myuploadfiles/");
    _upload_url = std::string("http://") + _host_ip +std::string(":8000/api/myv1_c/upload_single_definename");
    
    return true;
}

std::string TaskDealer::QueryTask_compile() const {
    std::string str;
    auto res = curl_get_req(_query_url.c_str(), str);
    return str;
}

std::string TaskDealer::UpdateTaskStatus(int id, const char* jsondata) const {
    std::string cmd = "curl -X PUT  -H \"Content-Type:application/json\" ";
    cmd += _update_url;
    cmd += std::to_string(id);
    cmd += " --data '";
    cmd += std::string(jsondata);
    cmd +="'";

    return curl_shell_put(cmd);
}

bool TaskDealer::Download_onnx(const char* filename) const {
    std::string url = _download_url + filename;
    return curl_download_file(url.c_str(), filename) == 0;
}

std::string TaskDealer::Upload_file(const char* local_path, const char* dest_path) const {
    std::string str1, str2;
    curl_upload_file(_upload_url.c_str(), local_path, dest_path, str1, &str2);
    return str2;
}
