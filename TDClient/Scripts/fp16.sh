# /opt/qti-aic/exec/qaic-exec -aic-num-cores=$1  -m=pointpillar.onnx  -aic-binary-dir=$2  \
# -convert-to-fp16 -aic-hw -aic-hw-version=2.0  -device-id=0 -v \
# -combine-outputs=true -version -aic-perf-metrics -use-producer-dma  \
# -aic-enable-depth-first  -aic-depth-first-mem=10 \
# -num-iter=10 -compile-only
/opt/qti-aic/exec/qaic-exec -aic-num-cores=2  -aic-binary-dir=point_2_fp16_0915 -m=pointpillar.onnx \
-aic-hw -aic-hw-version=2.0  -device-id=1  -v  -num-iter=10 \
-use-producer-dma=true   -compile-only -multicast-weights \
-aic-enable-depth-first   \
-combine-inputs=false -combine-outputs=false \
-convert-to-fp16  -aic-preproc \
