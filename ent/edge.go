// Code generated by entc, DO NOT EDIT.

package ent

import "context"

func (a *Activation) User(ctx context.Context) ([]*User, error) {
	result, err := a.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryUser().All(ctx)
	}
	return result, err
}

func (a *Activation) Actor(ctx context.Context) ([]*User, error) {
	result, err := a.Edges.ActorOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryActor().All(ctx)
	}
	return result, err
}

func (ml *MapLayer) Metadata(ctx context.Context) (*Metadata, error) {
	result, err := ml.Edges.MetadataOrErr()
	if IsNotLoaded(err) {
		result, err = ml.QueryMetadata().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ml *MapLayer) MapType(ctx context.Context) (*MapType, error) {
	result, err := ml.Edges.MapTypeOrErr()
	if IsNotLoaded(err) {
		result, err = ml.QueryMapType().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (mt *MapType) Metadata(ctx context.Context) (*Metadata, error) {
	result, err := mt.Edges.MetadataOrErr()
	if IsNotLoaded(err) {
		result, err = mt.QueryMetadata().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (mt *MapType) MapLayers(ctx context.Context) ([]*MapLayer, error) {
	result, err := mt.Edges.MapLayersOrErr()
	if IsNotLoaded(err) {
		result, err = mt.QueryMapLayers().All(ctx)
	}
	return result, err
}

func (mt *MapType) Servers(ctx context.Context) ([]*Server, error) {
	result, err := mt.Edges.ServersOrErr()
	if IsNotLoaded(err) {
		result, err = mt.QueryServers().All(ctx)
	}
	return result, err
}

func (m *Metadata) Schema(ctx context.Context) (*MetadataSchema, error) {
	result, err := m.Edges.SchemaOrErr()
	if IsNotLoaded(err) {
		result, err = m.QuerySchema().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (m *Metadata) User(ctx context.Context) (*User, error) {
	result, err := m.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (m *Metadata) MapType(ctx context.Context) (*MapType, error) {
	result, err := m.Edges.MapTypeOrErr()
	if IsNotLoaded(err) {
		result, err = m.QueryMapType().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ms *MetadataSchema) Metadata(ctx context.Context) ([]*Metadata, error) {
	result, err := ms.Edges.MetadataOrErr()
	if IsNotLoaded(err) {
		result, err = ms.QueryMetadata().All(ctx)
	}
	return result, err
}

func (pl *Player) Metadata(ctx context.Context) (*Metadata, error) {
	result, err := pl.Edges.MetadataOrErr()
	if IsNotLoaded(err) {
		result, err = pl.QueryMetadata().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pl *Player) Servers(ctx context.Context) ([]*Server, error) {
	result, err := pl.Edges.ServersOrErr()
	if IsNotLoaded(err) {
		result, err = pl.QueryServers().All(ctx)
	}
	return result, err
}

func (pl *Player) User(ctx context.Context) (*User, error) {
	result, err := pl.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = pl.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pl *Player) Identifiers(ctx context.Context) ([]*PlayerIdentifier, error) {
	result, err := pl.Edges.IdentifiersOrErr()
	if IsNotLoaded(err) {
		result, err = pl.QueryIdentifiers().All(ctx)
	}
	return result, err
}

func (pi *PlayerIdentifier) Metadata(ctx context.Context) (*Metadata, error) {
	result, err := pi.Edges.MetadataOrErr()
	if IsNotLoaded(err) {
		result, err = pi.QueryMetadata().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pi *PlayerIdentifier) Player(ctx context.Context) (*Player, error) {
	result, err := pi.Edges.PlayerOrErr()
	if IsNotLoaded(err) {
		result, err = pi.QueryPlayer().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (s *Server) Metadata(ctx context.Context) (*Metadata, error) {
	result, err := s.Edges.MetadataOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryMetadata().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (s *Server) ServerType(ctx context.Context) (*ServerType, error) {
	result, err := s.Edges.ServerTypeOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryServerType().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (s *Server) MapType(ctx context.Context) (*MapType, error) {
	result, err := s.Edges.MapTypeOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryMapType().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (s *Server) Players(ctx context.Context) ([]*Player, error) {
	result, err := s.Edges.PlayersOrErr()
	if IsNotLoaded(err) {
		result, err = s.QueryPlayers().All(ctx)
	}
	return result, err
}

func (st *ServerType) Servers(ctx context.Context) ([]*Server, error) {
	result, err := st.Edges.ServersOrErr()
	if IsNotLoaded(err) {
		result, err = st.QueryServers().All(ctx)
	}
	return result, err
}

func (st *SessionToken) User(ctx context.Context) (*User, error) {
	result, err := st.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = st.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Metadata(ctx context.Context) (*Metadata, error) {
	result, err := u.Edges.MetadataOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryMetadata().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) SessionTokens(ctx context.Context) ([]*SessionToken, error) {
	result, err := u.Edges.SessionTokensOrErr()
	if IsNotLoaded(err) {
		result, err = u.QuerySessionTokens().All(ctx)
	}
	return result, err
}

func (u *User) Activation(ctx context.Context) (*Activation, error) {
	result, err := u.Edges.ActivationOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryActivation().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Activations(ctx context.Context) ([]*Activation, error) {
	result, err := u.Edges.ActivationsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryActivations().All(ctx)
	}
	return result, err
}

func (u *User) Players(ctx context.Context) ([]*Player, error) {
	result, err := u.Edges.PlayersOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryPlayers().All(ctx)
	}
	return result, err
}