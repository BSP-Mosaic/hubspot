package hubspot

import "fmt"

const (
	crmBasePath = "crm"

	objectsBasePath = "objects"
)

type CRM struct {
	Company  CompanyService
	Contact  ContactService
	Deal     DealService
	Owner    OwnerService
	Pipeline PipelineService
}

func newCRM(c *Client) *CRM {
	crmPath := fmt.Sprintf("%s/%s", crmBasePath, c.apiVersion)
	return &CRM{
		Company: &CompanyServiceOp{
			companyPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, companyBasePath),
			client:      c,
		},
		Contact: &ContactServiceOp{
			contactPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, contactBasePath),
			client:      c,
		},
		Deal: &DealServiceOp{
			dealPath: fmt.Sprintf("%s/%s/%s", crmPath, objectsBasePath, dealBasePath),
			client:   c,
		},
		Owner: &OwnerServiceOp{
			ownerPath: fmt.Sprintf("%s/%s", crmPath, ownerBasePath),
			client:    c,
		},
		Pipeline: &PipelineServiceOp{
			pipelinePath: fmt.Sprintf("%s/%s", crmPath, pipelineBasePath),
			client:       c,
		},
	}
}
