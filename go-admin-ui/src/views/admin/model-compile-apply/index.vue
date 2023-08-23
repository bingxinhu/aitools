
<template>
    <BasicLayout>
        <template #wrapper>
            <el-card class="box-card">
                <el-form ref="queryForm" :model="queryParams" :inline="true" label-width="68px">
                    <el-form-item label="申请人" prop="applicant"><el-input v-model="queryParams.applicant" placeholder="请输入申请人" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="状态" prop="status"><el-input v-model="queryParams.status" placeholder="请输入状态" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="模型名称" prop="name"><el-input v-model="queryParams.name" placeholder="请输入模型名称" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="模型类别" prop="type"><el-input v-model="queryParams.type" placeholder="请输入模型类别" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        <el-form-item label="模型md5" prop="md5"><el-input v-model="queryParams.md5" placeholder="请输入模型md5" clearable
                                              size="small" @keyup.enter.native="handleQuery"/>
                            </el-form-item>
                        
                    <el-form-item>
                        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
                        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
                    </el-form-item>
                </el-form>

                <el-row :gutter="10" class="mb8">
                    <el-col :span="1.5">
                        <el-button
                                v-permisaction="['admin:modelCompileApply:add']"
                                type="primary"
                                icon="el-icon-plus"
                                size="mini"
                                @click="handleAdd"
                        >新增
                        </el-button>
                    </el-col>
                    <el-col :span="1.5" v-if="user_role == 'admin'||user_role == 'checker'">
                        <el-button
                                v-permisaction="['admin:modelCompileApply:edit']"
                                type="success"
                                icon="el-icon-edit"
                                size="mini"
                                :disabled="single"
                                @click="handleUpdate"
                        >修改
                        </el-button>
                    </el-col>
                    <el-col :span="1.5" v-if="user_role == 'admin'||user_role == 'checker'">
                        <el-button
                                v-permisaction="['admin:modelCompileApply:remove']"
                                type="danger"
                                icon="el-icon-delete"
                                size="mini"
                                :disabled="multiple"
                                @click="handleDelete"
                        >删除
                        </el-button>
                    </el-col>
                </el-row>

                <el-table v-loading="loading" :data="modelCompileApplyList" @selection-change="handleSelectionChange">
                    <el-table-column type="selection" width="55" align="center"/><el-table-column label="申请时间" align="center" prop="applyTime"
                                                 :show-overflow-tooltip="true">
                                    <template slot-scope="scope">
                                    <span>{{ parseTime(scope.row.applyTime) }}</span>
                                    </template>
                                </el-table-column><el-table-column label="申请人" align="center" prop="applicant"
                                                 :show-overflow-tooltip="true"/><el-table-column label="状态" align="center" prop="status"
                                                 :show-overflow-tooltip="true">
                                                 <template slot-scope="scope" >
                                                    {{ scope.row.status }}
                                                    <a target="_blank" style="color:#007bff;" v-if="scope.row.status=='编译完成'"
                                                    v-bind:href="'http://10.40.28.222:8000/static/myuploadfiles/'+scope.row.onnx+'_out.bin'">
                                                    下载
                                                    </a>
                                                <span v-else> {{ scope.row.goid1 }}</span>
                                                </template>
                                                 </el-table-column>
                                                 <el-table-column label="模型名称" align="center" prop="name"
                                                 :show-overflow-tooltip="true"/><el-table-column label="模型类别" align="center" prop="type"
                                                 :show-overflow-tooltip="true"/><el-table-column label="ONNX" align="center" prop="onnx_local"
                                                 :show-overflow-tooltip="true"> 
                                                 <template slot-scope="scope">
                                                    <a target="_blank" style="color:#007bff;" v-if="scope.row.onnx!=null"
                                                    v-bind:href="'http://10.40.28.222:8000/static/myuploadfiles/'+scope.row.onnx">
                                                        {{ scope.row.onnx_local }}
                                                    </a>
                                                <span v-else> {{ scope.row.goid1 }}</span>
                                                </template>
                                                 </el-table-column><el-table-column label="模型md5" align="center" prop="md5"
                                                 :show-overflow-tooltip="true"/><el-table-column label="NSP核数" align="center" prop="nspCnt"
                                                 :show-overflow-tooltip="true"/><el-table-column label="编译类型" align="center" prop="compileType"
                                                 :show-overflow-tooltip="true"/><el-table-column label="通道顺序" align="center" prop="channelOrder"
                                                 :show-overflow-tooltip="true"/>
                                                <!--
                                                 :show-overflow-tooltip="true"/><el-table-column label="量化数据集" align="center" prop="calibDataset"
                                                 :show-overflow-tooltip="true"/><el-table-column label="后处理代码" align="center" prop="postCode"-->
                    <el-table-column label="操作" align="center" class-name="small-padding fixed-width" v-if="user_role == 'admin'||user_role == 'checker'">
                        <template slot-scope="scope">
                         <el-popconfirm
                           class="delete-popconfirm"
                           title="确认要修改吗?"
                           confirm-button-text="修改"
                           @onConfirm="handleUpdate(scope.row)"
                         >
                           <el-button
                             slot="reference"
                             v-permisaction="['admin:modelCompileApply:edit']"
                             size="mini"
                             type="text"
                             icon="el-icon-edit"
                           >修改
                           </el-button>
                         </el-popconfirm>
                         <el-popconfirm
                            class="delete-popconfirm"
                            title="确认要删除吗?"
                            confirm-button-text="删除"
                            @onConfirm="handleDelete(scope.row)"
                         >
                            <el-button
                              slot="reference"
                              v-permisaction="['admin:modelCompileApply:remove']"
                              size="mini"
                              type="text"
                              icon="el-icon-delete"
                            >删除
                            </el-button>
                         </el-popconfirm>
                        </template>
                    </el-table-column>
                </el-table>

                <pagination
                        v-show="total>0"
                        :total="total"
                        :page.sync="queryParams.pageIndex"
                        :limit.sync="queryParams.pageSize"
                        @pagination="getList"
                />

                <!-- 添加或修改对话框 -->
                <el-dialog :title="title" :visible.sync="open" width="500px">
                    <el-form ref="form" :model="form" :rules="rules" label-width="80px">
                        
                                    <el-form-item label="目标平台" prop="platform">
                                        <el-select v-model="form.platform" placeholder="请选择" clearable :style="{width: '100%'}">
                                            <el-option v-for="(item, index) in platforms" :key="index" :label="item.label"
                                            :value="item.value" :disabled="item.disabled"></el-option>
                                        </el-select>
                                    </el-form-item>
                                    <el-form-item label="申请时间" prop="applyTime">
                                        <el-date-picker
                                                    v-model="form.applyTime"
                                                    type="datetime"
                                                    placeholder="选择日期">
                                            </el-date-picker>
                                    </el-form-item>
                                    <el-form-item label="申请人" prop="applicant">
                                        <el-input v-model="form.applicant" placeholder="申请人"
                                                      />
                                    </el-form-item>
                                    <!--el-form-item label="状态" prop="status" v-show="false">
                                        <el-input v-model="form.status" placeholder="状态"
                                                      />
                                    </el-form-item-->
                                    <el-form-item label="状态" prop="status"  v-if="user_role == 'admin'||user_role == 'checker'">
                                        <el-select v-model="form.status" placeholder="请选择" clearable :style="{width: '100%'}">
                                            <el-option v-for="(item, index) in statusOptions" :key="index" :label="item.label"
                                            :value="item.value" :disabled="item.disabled"></el-option>
                                        </el-select>
                                    </el-form-item>
                                    <el-form-item label="模型名称" prop="name">
                                        <el-input v-model="form.name" placeholder="模型名称"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="模型类别" prop="type">
                                        <el-input v-model="form.type" placeholder="模型类别"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="ONNX" prop="onnx">
                                        <el-row :gutter="10" class="mb8">
                                            <el-col :span="1.5">
                                                <el-input v-model="form.onnx" v-show="false"/>
                                                <el-input v-model="form.onnx_local" :disabled="true" placeholder="ONNX"/>
                                            </el-col>

                                            <el-col :span="1.5">
                                                <input 
                                                    type="file" 
                                                    name="avatar" 
                                                    accept="*.onnx" 
                                                    style="display:none"
                                                    @change="changeImage($event)" 
                                                    ref="avatarInput" />
                                                <el-button @click="handleSubmit">上传</el-button>
                                            </el-col>
                                        </el-row>
                                    </el-form-item>
                                    <el-form-item label="模型md5" prop="md5">
                                        <el-input v-model="form.md5" placeholder="模型md5"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="NSP核数" prop="nspCnt">
                                        <el-input v-model="form.nspCnt" placeholder="NSP核数"
                                                      />
                                    </el-form-item>
                                    <!--el-form-item label="编译类型" prop="compileType">
                                        <el-input v-model="form.compileType" placeholder="编译类型"
                                                      />
                                    </el-form-item-->
                                    <el-form-item label="编译类型" prop="compileType">
                                        <el-select v-model="form.compileType" placeholder="请选择" clearable :style="{width: '100%'}">
                                            <el-option v-for="(item, index) in field101Options" :key="index" :label="item.label"
                                            :value="item.value" :disabled="item.disabled"></el-option>
                                        </el-select>
                                    </el-form-item>

                                    <!--el-form-item label="通道顺序" prop="channelOrder">
                                        <el-select v-model="form.channelOrder" placeholder="请选择" clearable :style="{width: '100%'}">
                                            <el-option v-for="(item, index) in coOptions" :key="index" :label="item.label"
                                            :value="item.value" :disabled="item.disabled"></el-option>
                                        </el-select>
                                    </el-form-item-->
                                    <el-form-item label="通道顺序" prop="channelOrder">
                                        <el-radio-group v-model="form.channelOrder">
                                                <el-radio
                                                        v-for="dict in channelOptions"
                                                        :key="dict.value"
                                                        :label="dict.value"
                                                >{{ dict.label }}</el-radio>
                                            </el-radio-group>
                                    </el-form-item>
                                    <el-form-item label="量化数据集" prop="calibDataset">
                                        <el-input v-model="form.calibDataset" placeholder="量化数据集"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="后处理代码" prop="postCode">
                                        <el-input v-model="form.postCode" placeholder="后处理代码"
                                                      />
                                    </el-form-item>
                                    <el-form-item label="编译产出" prop="onnx_out" v-if="user_role == 'admin'||user_role == 'checker'">
                                        <el-row :gutter="10" class="mb8">
                                            <el-col :span="1.5">
                                                <el-input v-model="form.onnx_out" :disabled="true" placeholder="上传编译产出" v-show="false"/>
                                            </el-col>

                                            <el-col :span="1.5">
                                                <input 
                                                    type="file" 
                                                    name="avatar_out" 
                                                    accept="*.bin" 
                                                    style="display:none"
                                                    @change="changeImage_out($event)" 
                                                    ref="avatarInput_out" />
                                                <el-button @click="handleSubmit_out">上传</el-button>
                                            </el-col>
                                        </el-row>
                                    </el-form-item>
                    </el-form>
                    <div slot="footer" class="dialog-footer">
                        <el-button type="primary" @click="submitForm">确 定</el-button>
                        <el-button @click="cancel">取 消</el-button>
                    </div>
                </el-dialog>
            </el-card>
        </template>
    </BasicLayout>
</template>

<script>
    import {addModelCompileApply, delModelCompileApply, getModelCompileApply, listModelCompileApply, getUserRole, updateModelCompileApply, uploadFile, uploadFile_out} from '@/api/admin/model-compile-apply'
    
    export default {
        name: 'ModelCompileApply',
        components: {
        },
        data() {
            return {
                user_role: '',
                // 遮罩层
                loading: true,
                // 选中数组
                ids: [],
                // 非单个禁用
                single: true,
                // 非多个禁用
                multiple: true,
                // 总条数
                total: 0,
                // 弹出层标题
                title: '',
                // 是否显示弹出层
                open: false,
                isEdit: false,
                // 类型数据字典
                typeOptions: [],
                modelCompileApplyList: [],
                channelOptions: [{label:"NCHW", value:"NCHW"}, {label:"NHWC", value:"NHWC"}],
                
                // 关系表类型
                
                // 查询参数
                queryParams: {
                    pageIndex: 1,
                    pageSize: 10,
                    applicant:undefined,
                    status:undefined,
                    name:undefined,
                    type:undefined,
                    md5:undefined,
                    
                },
                // 表单参数
                form: {
                },
                field101Options: [{
                    "label": "INT8",
                    "value": "INT8"
                }, {
                    "label": "FINT8",
                    "value": "FINT8"
                }, {
                    "label": "FP16",
                    "value": "FP16"
                }],
                coOptions: [{
                    "label": "NCHW",
                    "value": "NCHW"
                }, {
                    "label": "NHWC",
                    "value": "NHWC"
                }],
                platforms: [{
                    "label": "SA9000",
                    "value": "SA9000"
                }, {
                    "label": "SA8540",
                    "value": "SA8540"
                }, {
                    "label": "TDA4",
                    "value": "TDA4"
                }, {
                    "label": "V4H",
                    "value": "V4H"
                }],
                statusOptions: [{
                    "label": "待审核",
                    "value": "待审核"
                }, {
                    "label": "待编译",
                    "value": "待编译"
                }, {
                    "label": "驳回",
                    "value": "驳回"
                }, {
                    "label": "人工编译",
                    "value": "人工编译"
                }, {
                    "label": "编译完成",
                    "value": "编译完成"
                }
                // , {
                //     "label": "编译中",
                //     "value": "编译中"
                // }, {
                //     "label": "编译完成",
                //     "value": "编译完成"
                // }, {
                //     "label": "编译失败",
                //     "value": "编译失败"
                // }
                ],
                // 表单校验
                rules: {applicant:  [ {required: true, message: '申请人不能为空', trigger: 'blur'} ],
                platform:  [ {required: true, message: '目标平台不能为空', trigger: 'blur'},{pattern: /^(SA9000)$/,message: "此平台尚未支持",trigger: "change"} ],
                status:  [ {required: true, message: '状态不能为空', trigger: 'blur'} ],
                name:  [ {required: true, message: '模型名称不能为空', trigger: 'blur'} ],
                type:  [ {required: true, message: '模型类别不能为空', trigger: 'blur'} ],
                onnx:  [ {required: true, message: 'onnx不能为空', trigger: 'blur'} ],
                md5:  [ {required: true, message: '模型md5不能为空', trigger: 'blur'} ],
                compileType:  [ {required: true, message: '编译类型不能为空', trigger: 'blur'} ],
                channelOrder:  [ {required: true, message: '通道顺序不能为空', trigger: 'blur'} ],
                nspCnt:  [ {required: true, message: 'NSP_num不能为空', trigger: 'blur'}, {pattern: /^([1-9]|1[0-4])$/,message: "NSP:1~14",trigger: "blur"}],
                }
        }
        },
        created() {
            this.getRole()
            this.getList()
            },
        methods: {
            getRole() {
                getUserRole().then(response => (this.user_role = response.data))
            },
            /** 查询参数列表 */
            getList() {
                this.loading = true
                listModelCompileApply(this.addDateRange(this.queryParams, this.dateRange)).then(response => {
                        this.modelCompileApplyList = response.data.list
                        this.total = response.data.count
                        this.loading = false
                    }
                )
            },
            // 取消按钮
            cancel() {
                this.open = false
                this.reset()
            },
            // 表单重置
            reset() {
                this.form = {
                
                id: undefined,
                applyTime: undefined,
                applicant: undefined,
                name: undefined,
                type: undefined,
                onnx: undefined,
                onnx_out: undefined,
                md5: undefined,
                nspCnt: undefined,
                compileType: undefined,
                calibDataset: undefined,
                postCode: undefined,
            }
                this.resetForm('form')
            },
            getImgList: function() {
              this.form[this.fileIndex] = this.$refs['fileChoose'].resultList[0].fullUrl
            },
            fileClose: function() {
              this.fileOpen = false
            },
            // 关系
            // 文件
            /** 搜索按钮操作 */
            handleQuery() {
                this.queryParams.pageIndex = 1
                this.getList()
            },
            /** 重置按钮操作 */
            resetQuery() {
                this.dateRange = []
                this.resetForm('queryForm')
                this.handleQuery()
            },
            /** 新增按钮操作 */
            handleAdd() {
                this.reset()
                this.open = true
                this.title = '添加模型编译申请'
                this.isEdit = false
                this.form.status = "待审核"
            },
            // 多选框选中数据
            handleSelectionChange(selection) {
                this.ids = selection.map(item => item.id)
                this.single = selection.length !== 1
                this.multiple = !selection.length
            },
            /** 修改按钮操作 */
            handleUpdate(row) {
                this.reset()
                const id =
                row.id || this.ids
                getModelCompileApply(id).then(response => {
                    this.form = response.data
                    this.open = true
                    this.title = '修改模型编译申请'
                    this.isEdit = true
                })
            },
            /** 提交按钮 */
            submitForm: function () {
                this.$refs['form'].validate(valid => {
                    if (valid) {
                        if (this.form.id !== undefined) {
                            updateModelCompileApply(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        } else {
                            addModelCompileApply(this.form).then(response => {
                                if (response.code === 200) {
                                    this.msgSuccess(response.msg)
                                    this.open = false
                                    this.getList()
                                } else {
                                    this.msgError(response.msg)
                                }
                            })
                        }
                    }
                })
            },
            /** 删除按钮操作 */
            handleDelete(row) {
                var Ids = (row.id && [row.id]) || this.ids

                this.$confirm('是否确认删除编号为"' + Ids + '"的数据项?', '警告', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(function () {
                      return delModelCompileApply( { 'ids': Ids })
                }).then((response) => {
                   if (response.code === 200) {
                     this.msgSuccess(response.msg)
                     this.open = false
                     this.getList()
                   } else {
                     this.msgError(response.msg)
                   }
                }).catch(function () {
                })
            },
            handleSubmit() {
                this.$refs.avatarInput.click()
            },
            // 当文件域的内容改变的时候获取里面的文件信息对象
            changeImage(e) {
                let that = this
                const file = e.target.files[0]
                var formData = new FormData()
                formData.append("file",file);
                uploadFile(formData).then(response => {
                                if (response.data === "null") {
                                    alert("upload failed!")
                                } else {
                                    this.form.onnx = response.data
                                    this.form.onnx_local = file.name
                                    //this.form.status = "待审核"
                                }
                            })
            },
            handleSubmit_out() {
                this.$refs.avatarInput_out.click()
            },
            // 当文件域的内容改变的时候获取里面的文件信息对象
            changeImage_out(e) {
                let that = this
                const file = e.target.files[0]
                var formData = new FormData()
                formData.append("file",file);
                formData.append("savename",this.form.onnx+"_out.bin");
                uploadFile_out(formData).then(response => {
                                if (response.data === "null") {
                                    alert("upload failed!")
                                } else {
                                    alert("上传成功")
                                    //this.form.onnx = response.data
                                    //this.form.onnx = "上传成功"
                                    //this.form.status = "待审核"
                                }
                            })
            }
        }
    }
</script>
