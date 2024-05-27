## 流控规则
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
