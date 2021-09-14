package rds

import (
	"github.com/aquasecurity/defsec/provider"
	"github.com/aquasecurity/defsec/rules"
	"github.com/aquasecurity/defsec/severity"
	"github.com/aquasecurity/defsec/state"
)

var CheckNoPublicDbAccess = rules.Register(
	rules.Rule{
		Provider:    provider.AWSProvider,
		Service:     "rds",
		ShortCode:   "no-public-db-access",
		Summary:     "A database resource is marked as publicly accessible.",
		Impact:      "The database instance is publicly accessible",
		Resolution:  "Set the database to not be publicly accessible",
		Explanation: `Database resources should not publicly available. You should limit all access to the minimum that is required for your application to function.`,
		Links:       []string{},
		Severity:    severity.Critical,
	},
	func(s *state.State) (results rules.Results) {
		for _, cluster := range s.AWS.RDS.Clusters {
			for _, instance := range cluster.Instances {
				if instance.PublicAccess.IsTrue() {
					results.Add(
						"Cluster instance is exposed publicly.",
						instance.PublicAccess.Metadata(),
						instance.PublicAccess.Value(),
					)
				}
			}
		}
		for _, instance := range s.AWS.RDS.Instances {
			if instance.PublicAccess.IsTrue() {
				results.Add(
					"Instance is exposed publicly.",
					instance.PublicAccess.Metadata(),
					instance.PublicAccess.Value(),
				)
			}
		}
		return
	},
)
