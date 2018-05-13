package basicnode

import (
	"sync"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/the-anna-project/the-anna-project/action"
	"github.com/the-anna-project/the-anna-project/network"
	"github.com/the-anna-project/the-anna-project/port"
	"github.com/the-anna-project/the-anna-project/random"
)

type Config struct {
	// Action is the business logic implementation the node executes when being
	// activated. Nodes might be configured with a no-op action when they only
	// serve signal dispatching purposes within the network graphs.
	Action  action.Interface
	Logger  micrologger.Logger
	Network network.Interface
	Random  random.Interface
}

type Object struct {
	action  action.Interface
	logger  micrologger.Logger
	network network.Interface
	random  random.Interface

	alreadyBooted   bool
	alreadyShutDown bool
	energy          float64
	id              string
	mutex           sync.Mutex
	inputPorts      []port.Interface
	outputPorts     []port.Interface
	shutdown        chan struct{}
	threshold       float64
}

func New(config Config) (*Object, error) {
	if config.Action == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Action must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.Network == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Network must not be empty", config)
	}
	if config.Random == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Random must not be empty", config)
	}

	o := &Object{
		action:  config.Action,
		logger:  config.Logger,
		network: config.Network,
		random:  config.Random,

		alreadyBooted:   false,
		alreadyShutDown: false,
		energy:          0,
		id:              "",
		mutex:           sync.Mutex{},
		inputPorts:      nil,
		outputPorts:     nil,
		shutdown:        make(chan struct{}, 1),
		threshold:       0,
	}

	return o, nil
}
