# GPT_Cli

## 项目框架

![project_graph](/pic/index.png)

## 部署运行

```shell
docker build -t your-image-name .
docker-compose up
```

镜像启动后暴露50051端口，服务会监听并处理50051端口数据，直接使用grpc连接并发送即可
## features

### gRPC 高吞吐收发

```go
func (s *server) SendString(ctx context.Context, msg *pb.StringMessage) (*pb.EmptyMessage, error) {
    // here, we send the instructions to GPT
    s.processedString = SendContent(msg.Content)
    log.Printf("ASR message sent: %s", msg.Content)
    return &pb.EmptyMessage{}, nil
}
```

### 完备可选token参数

```go

type ChatCompletionRequest struct {
    Model ChatGPTModel 
    Messages []ChatMessage
    Temperature float64 
    TopP float64
    N int
    MaxTokens int
    PresencePenalty float64
    FrequencyPenalty float64
    User string
}
```

### 标准化错误处理

```go
var (
    ErrAPIKeyRequired = errors.New("API Key is required")

    // ErrInvalidModel 输入的Model有误
    ErrInvalidModel = errors.New("invalid model")

    // ErrNoMessages 输入Message为空
    ErrNoMessages = errors.New("no messages provided")

    // ErrInvalidRole Role 仅支持user,system,assistant
    ErrInvalidRole = errors.New("invalid role. Only `user`, `system` and `assistant` are supported")

    // ErrInvalidTemperature Temperature，temp应在区间 [0, 2]
    ErrInvalidTemperature = errors.New("invalid temperature. 0 <= temp <= 2")

// ErrInvalidPresencePenalty presence penalty，应在区间 [-2, 2]
    ErrInvalidPresencePenalty = errors.New("invalid presence penalty. -2<= presence penalty <= 2")

// ErrInvalidFrequencyPenalty frequency penalty，应在区间 [-2, 2]
    ErrInvalidFrequencyPenalty = errors.New("invalid frequency penalty. -2<= frequency penalty <= 2")
)
```