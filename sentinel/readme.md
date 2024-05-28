## 1.流控规则
```go
type Rule struct {
	ID                     string                 `json:"id,omitempty"`
	Resource               string                 `json:"resource"`
	TokenCalculateStrategy TokenCalculateStrategy `json:"tokenCalculateStrategy"`
	ControlBehavior        ControlBehavior        `json:"controlBehavior"`
	Threshold              float64                `json:"threshold"`
	RelationStrategy       RelationStrategy       `json:"relationStrategy"`
	RefResource            string                 `json:"refResource"`
	MaxQueueingTimeMs      uint32                 `json:"maxQueueingTimeMs"`
	WarmUpPeriodSec        uint32                 `json:"warmUpPeriodSec"`
	WarmUpColdFactor       uint32                 `json:"warmUpColdFactor"`
	StatIntervalInMs       uint32                 `json:"statIntervalInMs"`
}
```

- Resource 资源名
- TokenCalculateStrategy 当前流量控制器的Token计算策略
  - Direct 直接使用字段 Threshold 作为阈值
  - WarmUp 使用预热方式计算Token的阈值
- ControlBehavior 流量控制器的控制策略
  - Reject 超过阈值直接拒绝
  - Throttling 匀速排队
- Threshold 流控阈值
  - 如果StatIntervalInMs是1000，那么Threshold就表示QPS，流量控制器也就会依据资源的QPS来做流控
- RelationStrategy 调用关系限流策略
  - CurrentResource 当前规则的resource做流控
  - AssociatedResource 使用关联的resource做流控
- RefResource 关联的resource
- WarmUpPeriodSec 预热的时间长度
  - 该字段仅仅对WarmUp的TokenCalculateStrategy生效
- WarmUpColdFactor 预热的因子，默认为3
  - 该值的设置会影响预热的速度，该字段仅仅对WarmUp的TokenCalculateStrategy生效
- MaxQueueingTimeMs 匀速排队的最大等待时间
  - 该字段仅仅对Throttling的ControlBehavior生效
- StatIntervalInMs 规则对应的流量控制器的独立统计结构的统计周期
  - 如果为1000，就是统计QPS

## 熔断降级规则
```go
type Rule struct {
	Id               string   `json:"id,omitempty"`
	Resource         string   `json:"resource"`
	Strategy         Strategy `json:"strategy"`
	RetryTimeoutMs   uint32   `json:"retryTimeoutMs"`
	MinRequestAmount uint64   `json:"minRequestAmount"`
	StatIntervalMs   uint32   `json:"statIntervalMs"`
	MaxAllowedRtMs   uint64   `json:"maxAllowedRtMs"`
	Threshold        float64  `json:"threshold"`
}
```

- Id 规则的全局唯一Id
- Resoure 熔断器规则生效的埋点资源的名称
- Strategy 熔断策略，目前支持SlowRequestRatio、ErrorRatio、ErrorCount三种
  - SlowRequestRatio 慢调用比例，需要设置允许的最大响应时间MaxAllowedRtMs，请求的响应时间大于该值则统计为慢调用。单位统计时长内，慢调用值大于Threshold字段设置的阈值，则触发熔断，经过熔断时长后熔断器会进入探测恢复阶段，若接下来一个请求响应时间小于设置的最大RT则结束熔断，若大于设置的RT则会再次被熔断
  - ErrorRatio 错误比例，触发与SlowRequestRatio类似，只不过判断标准为错误数，若接下来的一个请求没有错误则结束熔断，否则会再次熔断，代码中可以用 api.TraceError(entry, err) 来记录error数
- RetryTimeoutMs 熔断触发后持续时间
  - 资源进入熔断状态后，在配置的熔断时长内，请求都会快速失败。熔断结束之后进入探测恢复模式
- MinRequestAmount 静默数量
  - 如果当前统计周期内的请求数小于此值，即使达到熔断条件规则也不会触发
- StatIntervalMs 统计的时间窗口长度
- MaxAllowedRtMs 慢调用临界值
- Threshold 触发阈值
