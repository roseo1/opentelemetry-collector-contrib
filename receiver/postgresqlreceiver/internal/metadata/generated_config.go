// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for postgresqlreceiver metrics.
type MetricsConfig struct {
	PostgresqlBackends                 MetricConfig `mapstructure:"postgresql.backends"`
	PostgresqlBgwriterBuffersAllocated MetricConfig `mapstructure:"postgresql.bgwriter.buffers.allocated"`
	PostgresqlBgwriterBuffersWrites    MetricConfig `mapstructure:"postgresql.bgwriter.buffers.writes"`
	PostgresqlBgwriterCheckpointCount  MetricConfig `mapstructure:"postgresql.bgwriter.checkpoint.count"`
	PostgresqlBgwriterDuration         MetricConfig `mapstructure:"postgresql.bgwriter.duration"`
	PostgresqlBgwriterMaxwritten       MetricConfig `mapstructure:"postgresql.bgwriter.maxwritten"`
	PostgresqlBlocksRead               MetricConfig `mapstructure:"postgresql.blocks_read"`
	PostgresqlCommits                  MetricConfig `mapstructure:"postgresql.commits"`
	PostgresqlConnectionMax            MetricConfig `mapstructure:"postgresql.connection.max"`
	PostgresqlDatabaseCount            MetricConfig `mapstructure:"postgresql.database.count"`
	PostgresqlDbSize                   MetricConfig `mapstructure:"postgresql.db_size"`
	PostgresqlIndexScans               MetricConfig `mapstructure:"postgresql.index.scans"`
	PostgresqlIndexSize                MetricConfig `mapstructure:"postgresql.index.size"`
	PostgresqlOperations               MetricConfig `mapstructure:"postgresql.operations"`
	PostgresqlReplicationDataDelay     MetricConfig `mapstructure:"postgresql.replication.data_delay"`
	PostgresqlRollbacks                MetricConfig `mapstructure:"postgresql.rollbacks"`
	PostgresqlRows                     MetricConfig `mapstructure:"postgresql.rows"`
	PostgresqlTableCount               MetricConfig `mapstructure:"postgresql.table.count"`
	PostgresqlTableSize                MetricConfig `mapstructure:"postgresql.table.size"`
	PostgresqlTableVacuumCount         MetricConfig `mapstructure:"postgresql.table.vacuum.count"`
	PostgresqlWalAge                   MetricConfig `mapstructure:"postgresql.wal.age"`
	PostgresqlWalLag                   MetricConfig `mapstructure:"postgresql.wal.lag"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		PostgresqlBackends: MetricConfig{
			Enabled: true,
		},
		PostgresqlBgwriterBuffersAllocated: MetricConfig{
			Enabled: true,
		},
		PostgresqlBgwriterBuffersWrites: MetricConfig{
			Enabled: true,
		},
		PostgresqlBgwriterCheckpointCount: MetricConfig{
			Enabled: true,
		},
		PostgresqlBgwriterDuration: MetricConfig{
			Enabled: true,
		},
		PostgresqlBgwriterMaxwritten: MetricConfig{
			Enabled: true,
		},
		PostgresqlBlocksRead: MetricConfig{
			Enabled: true,
		},
		PostgresqlCommits: MetricConfig{
			Enabled: true,
		},
		PostgresqlConnectionMax: MetricConfig{
			Enabled: true,
		},
		PostgresqlDatabaseCount: MetricConfig{
			Enabled: true,
		},
		PostgresqlDbSize: MetricConfig{
			Enabled: true,
		},
		PostgresqlIndexScans: MetricConfig{
			Enabled: true,
		},
		PostgresqlIndexSize: MetricConfig{
			Enabled: true,
		},
		PostgresqlOperations: MetricConfig{
			Enabled: true,
		},
		PostgresqlReplicationDataDelay: MetricConfig{
			Enabled: true,
		},
		PostgresqlRollbacks: MetricConfig{
			Enabled: true,
		},
		PostgresqlRows: MetricConfig{
			Enabled: true,
		},
		PostgresqlTableCount: MetricConfig{
			Enabled: true,
		},
		PostgresqlTableSize: MetricConfig{
			Enabled: true,
		},
		PostgresqlTableVacuumCount: MetricConfig{
			Enabled: true,
		},
		PostgresqlWalAge: MetricConfig{
			Enabled: true,
		},
		PostgresqlWalLag: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesConfig provides config for postgresqlreceiver resource attributes.
type ResourceAttributesConfig struct {
	PostgresqlDatabaseName ResourceAttributeConfig `mapstructure:"postgresql.database.name"`
	PostgresqlIndexName    ResourceAttributeConfig `mapstructure:"postgresql.index.name"`
	PostgresqlTableName    ResourceAttributeConfig `mapstructure:"postgresql.table.name"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		PostgresqlDatabaseName: ResourceAttributeConfig{
			Enabled: true,
		},
		PostgresqlIndexName: ResourceAttributeConfig{
			Enabled: true,
		},
		PostgresqlTableName: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for postgresqlreceiver metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
