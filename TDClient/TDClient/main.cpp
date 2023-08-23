#include "cJSON.h"
#include "TaskDealer.h"
#include <unistd.h>

#include <stdlib.h>
//#include "curlfunc.hpp"

int main(int argc, char** argv) {
    TaskDealer dealer;
    if(!dealer.Init("192.168.0.209"))
        return -1;
    
    while(true) {
        usleep(5000000);
        std::string str = dealer.QueryTask_compile();
        cJSON* root = cJSON_Parse(str.c_str());
        if(!root) 
            continue;
        cJSON* item = cJSON_GetObjectItem(root, "data");
        if(!item)
            continue;
        item = cJSON_GetObjectItem(item, "list");
        if(!item)
            continue;
        
        if(cJSON_GetArraySize(item) < 1) {
            printf("No compile task!\n");

            cJSON_Delete(root);
            continue;
        }
        
        item = cJSON_GetArrayItem(item, 0);
        // get id
        int id = cJSON_GetObjectItem(item, "id")->valueint;
        std::string onnx_file = cJSON_GetObjectItem(item, "onnx")->valuestring;
        std::string nsp_cnt = cJSON_GetObjectItem(item, "nspCnt")->valuestring;
        std::string type = cJSON_GetObjectItem(item, "type")->valuestring;
        std::string compile_type = cJSON_GetObjectItem(item, "compileType")->valuestring;
        std::string channel_order = cJSON_GetObjectItem(item, "channelOrder")->valuestring;

        // change status
        cJSON_DeleteItemFromObject(item, "status");
        cJSON_AddStringToObject(item, "status", "编译中");
        char * itemstr = cJSON_Print(item);
        str = dealer.UpdateTaskStatus(id, itemstr);
        free(itemstr);
        cJSON* root_r = cJSON_Parse(str.c_str());
        if(!root_r)
            continue;
        cJSON* item_tmp = cJSON_GetObjectItem(root_r, "msg");
        if(!item_tmp || std::string(item_tmp->valuestring) != "修改成功")
            continue;
        cJSON_Delete(root_r);

        printf("Begin to compile...\n");

        // download onnx
        if(!dealer.Download_onnx(onnx_file.c_str())) {
            printf("onnx download failed!\n");
            continue;
        }
;
        // compile (block until finished this task)
        std::string compile_cmd = "../Scripts/compiler_sa9000.sh ./";//pointpillar.onnx MTN 2 FP16 NHWC";
        compile_cmd += onnx_file; compile_cmd+=" ";
        compile_cmd += type; compile_cmd+=" ";
        compile_cmd += nsp_cnt; compile_cmd+=" ";
        compile_cmd += compile_type; compile_cmd+=" ";
        compile_cmd += channel_order;

        printf("%s \n", compile_cmd.c_str());
        int ret = system(compile_cmd.c_str());
        printf("ret = %d ret\n", ret);
        if(ret == 0) {
            // upload compileoutputs 
            str = dealer.Upload_file((onnx_file + "_out/programqpc.bin").c_str(), (onnx_file + "_out.bin").c_str());
            printf("*%s\n", str.c_str());
        }

        //&& update status
        cJSON_DeleteItemFromObject(item, "status");
        cJSON_AddStringToObject(item, "status", ret==0?"编译完成":"编译失败");
        itemstr  = cJSON_Print(item);
        str = dealer.UpdateTaskStatus(id, itemstr);
        free(itemstr);
        root_r = cJSON_Parse(str.c_str());
        if(!root_r)
            continue;
        item_tmp = cJSON_GetObjectItem(root_r, "msg");
        if(!item_tmp || std::string(item_tmp->valuestring) != "修改成功")
            continue;
        cJSON_Delete(root_r);
       
        cJSON_Delete(root);
        //free(itemstr);
    }
    return 0;
}



// #include "curlfunc.hpp"
// #include "cJSON.h"

// using namespace std;
// //get请求和post请求数据响应函数
// int main()
// {
//     string getUrlStr = "http://10.40.29.152:8000/api/v1/model-compile-apply/query_noauth?status=待审核";
//     string getResponseStr;
//     auto res = curl_get_req(getUrlStr, getResponseStr);
//     if (res == CURLE_OK)
//     {
//         cout << getResponseStr << endl;
//     }


//     std::string cmd = "curl -X PUT  -H \"Content-Type:application/json\" http://10.40.29.152:8000/api/v1/model-compile-apply/update_noauth/16 --data '{\"id\":16,\"applyTime\":\"2022-11-14T15:28:20+08:00\",\"applicant\":\"刘方盛2\",\"status\":\"failed\",\"name\":\"啊啊啊\",\"type\":\"啊啊\",\"onnx\":\"e2b186d17a4ad4fed6a5ef62f36ab340.png\",\"onnx_local\":\"住宿2.png\",\"md5\":\"啊\",\"nspCnt\":\"2\",\"compileType\":\"INT8\",\"channelOrder\":\"NHWC\",\"calibDataset\":\"\",\"postCode\":\"\",\"createdAt\":\"2022-11-14T15:28:42+08:00\",\"updatedAt\":\"2022-11-14T15:36:05+08:00\",\"createBy\":1,\"updateBy\":1}'";
//     std::string result = curl_shell_put(cmd);
//     cout<<"result:"<<result<<endl;

//     cJSON*root=cJSON_Parse(result.c_str()); 
//     cJSON*item=cJSON_GetObjectItem(root,"msg"); 
//     if(item) {
//         cout << item->valuestring<<std::endl;
//     }


//     return 0;
// }


 //curl -X PUT  -H "Content-Type:application/json" http://10.40.29.152:8000/api/v1/model-compile-apply/update_noauth/16 --data '{"id":16,"applyTime":"2022-11-14T15:28:20+08:00","applicant":"刘方盛1","status":"failed","name":"啊啊啊","type":"啊啊","onnx":"e2b186d17a4ad4fed6a5ef62f36ab340.png","onnx_local":"住宿2.png","md5":"啊","nspCnt":"2","compileType":"INT8","channelOrder":"NHWC","calibDataset":"","postCode":"","createdAt":"2022-11-14T15:28:42+08:00","updatedAt":"2022-11-14T15:36:05+08:00","createBy":1,"updateBy":1}'
