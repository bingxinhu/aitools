#!/bin/bash
# 参数说明
# $1    onnx
# $2    type(MTN LANELOC TSR OBSTACLE...)
# $3    NSP num
# $4    compileType(FP16, FINT8, INT8)
# $5    channel (NHWC, NCHW)

# return 0:sucess 1: failed 2: param error
if [ $# != 5 ];then
    echo 'param num error'
    exit 2
fi

echo 'start compile, params:', $1, $2, $3, $4, $5
if [ "$4" == "FP16" ]; then
    echo 'FP16'
    tmppath=`basename $1`_out
    rm -rf $tmppath
    /opt/qti-aic/exec/qaic-exec -aic-num-cores=$3  -aic-binary-dir=$tmppath -m=$1 -aic-hw -aic-hw-version=2.0  -device-id=1  -v  -num-iter=10 -use-producer-dma=true   -compile-only -multicast-weights -aic-enable-depth-first  -combine-inputs=false -combine-outputs=false -convert-to-fp16  -aic-preproc
    exit $?
elif [ "$4" == 'IN8' ]; then
    echo 'IN8'
    exit 1
elif [ "$4" == 'FIN8' ]; then
    echo 'FIN8'
    exit 1
else
    exit 2
fi

# /opt/qti-aic/exec/qaic-exec -aic-num-cores=$1  -m=pointpillar.onnx  -aic-binary-dir=$2  \
# -convert-to-fp16 -aic-hw -aic-hw-version=2.0  -device-id=0 -v \
# -combine-outputs=true -version -aic-perf-metrics -use-producer-dma  \
# -aic-enable-depth-first  -aic-depth-first-mem=10 \
# -num-iter=10 -compile-only

# /opt/qti-aic/exec/qaic-exec -aic-num-cores=2  -aic-binary-dir=point_2_fp16_0915 -m=pointpillar.onnx \
# -aic-hw -aic-hw-version=2.0  -device-id=1  -v  -num-iter=10 \
# -use-producer-dma=true   -compile-only -multicast-weights \
# -aic-enable-depth-first   \
# -combine-inputs=false -combine-outputs=false \
# -convert-to-fp16  -aic-preproc \
