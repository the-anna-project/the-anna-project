package node

import (
	"sync"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/the-anna-project/the-anna-project/spec/action"
	"github.com/the-anna-project/the-anna-project/spec/network"
	"github.com/the-anna-project/the-anna-project/spec/peer"
	"github.com/the-anna-project/the-anna-project/spec/random"
)

type Config struct {
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
	peers           []peer.Interface
	shutdown        chan struct{}
	threshold       float64
}

func New(config Config) (*Object, error) {
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
		peers:           nil,
		shutdown:        make(chan struct{}, 1),
		threshold:       0,
	}

	return o, nil
}
