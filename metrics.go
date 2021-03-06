package centrifuge

import (
	"github.com/prometheus/client_golang/prometheus"
)

var metricsNamespace = "centrifuge"

var (
	messagesSentCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: "node",
		Name:      "messages_sent_count",
		Help:      "Number of messages sent.",
	}, []string{"type"})

	messagesReceivedCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: "node",
		Name:      "messages_received_count",
		Help:      "Number of messages received.",
	}, []string{"type"})

	actionCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: "node",
		Name:      "action_count",
		Help:      "Number of node actions called.",
	}, []string{"action"})

	numClientsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metricsNamespace,
		Subsystem: "node",
		Name:      "num_clients",
		Help:      "Number of clients connected.",
	})

	numUsersGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metricsNamespace,
		Subsystem: "node",
		Name:      "num_users",
		Help:      "Number of unique users connected.",
	})

	buildInfoGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: metricsNamespace,
		Subsystem: "node",
		Name:      "build",
		Help:      "Node build info.",
	}, []string{"version"})

	numChannelsGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: metricsNamespace,
		Subsystem: "node",
		Name:      "num_channels",
		Help:      "Number of channels with one or more subscribers.",
	})

	replyErrorCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: "client",
		Name:      "num_reply_errors",
		Help:      "Number of errors in replies sent to clients.",
	}, []string{"method"})

	commandDurationSummary = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  metricsNamespace,
		Subsystem:  "client",
		Name:       "command_duration_seconds",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		Help:       "Client command duration summary.",
	}, []string{"method"})

	apiHandlerDurationSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace:  metricsNamespace,
		Subsystem:  "http",
		Name:       "api_request_duration_seconds",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		Help:       "Duration of API handler in general.",
	})

	apiCommandDurationSummary = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  metricsNamespace,
		Subsystem:  "http",
		Name:       "api_request_command_duration_seconds",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		Help:       "Duration of API per command.",
	}, []string{"method"})

	transportConnectCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: "transport",
		Name:      "connect_count",
		Help:      "Number of connections to specific transport.",
	}, []string{"transport"})

	transportMessagesSent = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: metricsNamespace,
		Subsystem: "transport",
		Name:      "messages_sent",
		Help:      "Number of messages sent over specific transport.",
	}, []string{"transport"})
)

func init() {
	prometheus.MustRegister(messagesSentCount)
	prometheus.MustRegister(messagesReceivedCount)
	prometheus.MustRegister(actionCount)
	prometheus.MustRegister(numClientsGauge)
	prometheus.MustRegister(numUsersGauge)
	prometheus.MustRegister(numChannelsGauge)
	prometheus.MustRegister(commandDurationSummary)
	prometheus.MustRegister(replyErrorCount)
	prometheus.MustRegister(apiHandlerDurationSummary)
	prometheus.MustRegister(apiCommandDurationSummary)
	prometheus.MustRegister(transportConnectCount)
	prometheus.MustRegister(transportMessagesSent)
	prometheus.MustRegister(buildInfoGauge)
}
