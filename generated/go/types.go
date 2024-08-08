// This file was auto-generated by Fern from our API Definition.

package methodk8s

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/method-security/methodk8s/generated/go/core"
)

type ProtocolTypes string

const (
	ProtocolTypesHttp      ProtocolTypes = "HTTP"
	ProtocolTypesHttps     ProtocolTypes = "HTTPS"
	ProtocolTypesTcp       ProtocolTypes = "TCP"
	ProtocolTypesTls       ProtocolTypes = "TLS"
	ProtocolTypesUdp       ProtocolTypes = "UDP"
	ProtocolTypesTcpUdp    ProtocolTypes = "TCP_UDP"
	ProtocolTypesGeneve    ProtocolTypes = "GENEVE"
	ProtocolTypesUndefined ProtocolTypes = "UNDEFINED"
)

func NewProtocolTypesFromString(s string) (ProtocolTypes, error) {
	switch s {
	case "HTTP":
		return ProtocolTypesHttp, nil
	case "HTTPS":
		return ProtocolTypesHttps, nil
	case "TCP":
		return ProtocolTypesTcp, nil
	case "TLS":
		return ProtocolTypesTls, nil
	case "UDP":
		return ProtocolTypesUdp, nil
	case "TCP_UDP":
		return ProtocolTypesTcpUdp, nil
	case "GENEVE":
		return ProtocolTypesGeneve, nil
	case "UNDEFINED":
		return ProtocolTypesUndefined, nil
	}
	var t ProtocolTypes
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (p ProtocolTypes) Ptr() *ProtocolTypes {
	return &p
}

type Gateway struct {
	Name      string      `json:"name" url:"name"`
	Namespace string      `json:"namespace" url:"namespace"`
	Instance  string      `json:"instance" url:"instance"`
	Listeners []*Listener `json:"listeners,omitempty" url:"listeners,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *Gateway) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *Gateway) UnmarshalJSON(data []byte) error {
	type unmarshaler Gateway
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = Gateway(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *Gateway) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type Ingress struct {
	Name      string  `json:"name" url:"name"`
	Namespace string  `json:"namespace" url:"namespace"`
	Rules     []*Rule `json:"rules,omitempty" url:"rules,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (i *Ingress) GetExtraProperties() map[string]interface{} {
	return i.extraProperties
}

func (i *Ingress) UnmarshalJSON(data []byte) error {
	type unmarshaler Ingress
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = Ingress(value)

	extraProperties, err := core.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties

	i._rawJSON = json.RawMessage(data)
	return nil
}

func (i *Ingress) String() string {
	if len(i._rawJSON) > 0 {
		if value, err := core.StringifyJSON(i._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

type IngressReport struct {
	Gateways  []*Gateway `json:"gateways,omitempty" url:"gateways,omitempty"`
	Ingresses []*Ingress `json:"ingresses,omitempty" url:"ingresses,omitempty"`
	Errors    []string   `json:"Errors,omitempty" url:"Errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (i *IngressReport) GetExtraProperties() map[string]interface{} {
	return i.extraProperties
}

func (i *IngressReport) UnmarshalJSON(data []byte) error {
	type unmarshaler IngressReport
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = IngressReport(value)

	extraProperties, err := core.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties

	i._rawJSON = json.RawMessage(data)
	return nil
}

func (i *IngressReport) String() string {
	if len(i._rawJSON) > 0 {
		if value, err := core.StringifyJSON(i._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

type Listener struct {
	Name     string        `json:"name" url:"name"`
	Port     int           `json:"port" url:"port"`
	Protocol ProtocolTypes `json:"protocol" url:"protocol"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *Listener) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *Listener) UnmarshalJSON(data []byte) error {
	type unmarshaler Listener
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = Listener(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *Listener) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type Rule struct {
	Host        string `json:"host" url:"host"`
	Path        string `json:"path" url:"path"`
	ServiceName string `json:"serviceName" url:"serviceName"`
	ServicePort int    `json:"servicePort" url:"servicePort"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *Rule) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *Rule) UnmarshalJSON(data []byte) error {
	type unmarshaler Rule
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = Rule(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *Rule) String() string {
	if len(r._rawJSON) > 0 {
		if value, err := core.StringifyJSON(r._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

type Address struct {
	Type    string `json:"type" url:"type"`
	Address string `json:"address" url:"address"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (a *Address) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *Address) UnmarshalJSON(data []byte) error {
	type unmarshaler Address
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = Address(value)

	extraProperties, err := core.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties

	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *Address) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type Node struct {
	Name         string     `json:"name" url:"name"`
	Arch         string     `json:"arch" url:"arch"`
	Os           string     `json:"os" url:"os"`
	Instancetype string     `json:"instancetype" url:"instancetype"`
	Status       bool       `json:"status" url:"status"`
	Addresses    []*Address `json:"addresses,omitempty" url:"addresses,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (n *Node) GetExtraProperties() map[string]interface{} {
	return n.extraProperties
}

func (n *Node) UnmarshalJSON(data []byte) error {
	type unmarshaler Node
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = Node(value)

	extraProperties, err := core.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties

	n._rawJSON = json.RawMessage(data)
	return nil
}

func (n *Node) String() string {
	if len(n._rawJSON) > 0 {
		if value, err := core.StringifyJSON(n._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(n); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", n)
}

type NodeReport struct {
	Nodes      []*Node  `json:"nodes,omitempty" url:"nodes,omitempty"`
	ClusterUrl *string  `json:"clusterUrl,omitempty" url:"clusterUrl,omitempty"`
	Errors     []string `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (n *NodeReport) GetExtraProperties() map[string]interface{} {
	return n.extraProperties
}

func (n *NodeReport) UnmarshalJSON(data []byte) error {
	type unmarshaler NodeReport
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = NodeReport(value)

	extraProperties, err := core.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties

	n._rawJSON = json.RawMessage(data)
	return nil
}

func (n *NodeReport) String() string {
	if len(n._rawJSON) > 0 {
		if value, err := core.StringifyJSON(n._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(n); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", n)
}

type Container struct {
	Name            string           `json:"name" url:"name"`
	Image           string           `json:"image" url:"image"`
	Ports           []*ContainerPort `json:"ports,omitempty" url:"ports,omitempty"`
	SecurityContext *SecurityContext `json:"securityContext,omitempty" url:"securityContext,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *Container) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *Container) UnmarshalJSON(data []byte) error {
	type unmarshaler Container
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = Container(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *Container) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type ContainerPort struct {
	Port     int            `json:"port" url:"port"`
	Protocol *ProtocolTypes `json:"protocol,omitempty" url:"protocol,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *ContainerPort) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *ContainerPort) UnmarshalJSON(data []byte) error {
	type unmarshaler ContainerPort
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ContainerPort(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ContainerPort) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type Pod struct {
	Name       string       `json:"name" url:"name"`
	Namespace  *string      `json:"namespace,omitempty" url:"namespace,omitempty"`
	Version    *string      `json:"version,omitempty" url:"version,omitempty"`
	Node       *string      `json:"node,omitempty" url:"node,omitempty"`
	Status     *Status      `json:"status,omitempty" url:"status,omitempty"`
	Containers []*Container `json:"containers,omitempty" url:"containers,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *Pod) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *Pod) UnmarshalJSON(data []byte) error {
	type unmarshaler Pod
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = Pod(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *Pod) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PodReport struct {
	Pods       []*Pod   `json:"pods,omitempty" url:"pods,omitempty"`
	ClusterUrl *string  `json:"clusterUrl,omitempty" url:"clusterUrl,omitempty"`
	Errors     []string `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PodReport) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PodReport) UnmarshalJSON(data []byte) error {
	type unmarshaler PodReport
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PodReport(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PodReport) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type SecurityContext struct {
	RunAsRoot                *bool `json:"runAsRoot,omitempty" url:"runAsRoot,omitempty"`
	AllowPrivilegeEscalation *bool `json:"allowPrivilegeEscalation,omitempty" url:"allowPrivilegeEscalation,omitempty"`
	ReadOnlyRootFilesystem   *bool `json:"readOnlyRootFilesystem,omitempty" url:"readOnlyRootFilesystem,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SecurityContext) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SecurityContext) UnmarshalJSON(data []byte) error {
	type unmarshaler SecurityContext
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SecurityContext(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SecurityContext) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type Status struct {
	Status *StatusTypes `json:"status,omitempty" url:"status,omitempty"`
	PodIp  *string      `json:"podIp,omitempty" url:"podIp,omitempty"`
	HostIp *string      `json:"hostIp,omitempty" url:"hostIp,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *Status) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *Status) UnmarshalJSON(data []byte) error {
	type unmarshaler Status
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = Status(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *Status) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type StatusTypes string

const (
	StatusTypesPending   StatusTypes = "PENDING"
	StatusTypesRunning   StatusTypes = "RUNNING"
	StatusTypesSucceeded StatusTypes = "SUCCEEDED"
	StatusTypesFailed    StatusTypes = "FAILED"
	StatusTypesUnknown   StatusTypes = "UNKNOWN"
)

func NewStatusTypesFromString(s string) (StatusTypes, error) {
	switch s {
	case "PENDING":
		return StatusTypesPending, nil
	case "RUNNING":
		return StatusTypesRunning, nil
	case "SUCCEEDED":
		return StatusTypesSucceeded, nil
	case "FAILED":
		return StatusTypesFailed, nil
	case "UNKNOWN":
		return StatusTypesUnknown, nil
	}
	var t StatusTypes
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s StatusTypes) Ptr() *StatusTypes {
	return &s
}

type Service struct {
	Name      string         `json:"name" url:"name"`
	Namespace string         `json:"namespace" url:"namespace"`
	Type      string         `json:"type" url:"type"`
	ManagedBy string         `json:"managedBy" url:"managedBy"`
	Ports     []*ServicePort `json:"ports,omitempty" url:"ports,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *Service) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *Service) UnmarshalJSON(data []byte) error {
	type unmarshaler Service
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = Service(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *Service) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type ServicePort struct {
	Name       string        `json:"name" url:"name"`
	Protocol   ProtocolTypes `json:"protocol" url:"protocol"`
	Port       int           `json:"port" url:"port"`
	TargetPort string        `json:"targetPort" url:"targetPort"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *ServicePort) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *ServicePort) UnmarshalJSON(data []byte) error {
	type unmarshaler ServicePort
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = ServicePort(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *ServicePort) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type ServiceReport struct {
	Services []*Service `json:"services,omitempty" url:"services,omitempty"`
	Errors   []string   `json:"errors,omitempty" url:"errors,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *ServiceReport) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *ServiceReport) UnmarshalJSON(data []byte) error {
	type unmarshaler ServiceReport
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = ServiceReport(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *ServiceReport) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}
