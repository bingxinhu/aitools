#ifndef TASK_DEALER_H
#define TASK_DEALER_H

/***
 * Created by liufangsheng 
 * 2022.11.21
 * Class Used to Deal with Task produced by goadmin
*/

// 任务种类
enum TaskType {
    task_compile = 0,   // 编译模型
    task_infer,         // 执行推理

};

// 平台类型
enum PlatformType {
    platform_sa9000,
    platform_sa8540,
    platform_v4h,
    platform_nvidia,

};


#include <string>
/**
 * @brief The TaskDealer class 任务处理器类
 */
class TaskDealer {
public:
    TaskDealer(PlatformType platform=platform_sa9000);

    bool Init(const char* host_ip);

    std::string QueryTask_compile() const;
    std::string UpdateTaskStatus(int id, const char* jsondata) const;
    bool Download_onnx(const char* filename) const;
    std::string Upload_file(const char* local_path, const char* dest_path) const;
private:
    PlatformType _platform;
    std::string _host_ip;
    std::string _query_url;
    std::string _update_url;    // +id --data 'xxxxxx'(json data)
    std::string _download_url;
    std::string _upload_url;
};

#endif//TASK_DEALER_H
