package google

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"google.golang.org/api/siteverification/v1"
	"log"
)

func resourceSiteVerification() *schema.Resource {
	return &schema.Resource{
		Create: resourceSiteVerificationCreate,
		Read:   resourceSiteVerificationRead,
		Delete: resourceSiteVerificationDelete,

		Schema: map[string]*schema.Schema{
			"identifier": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain name to verfiy",
				ForceNew:    true,
			},
			"method": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: siteVerificationMethodValid,
			},
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceSiteVerificationCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	method := d.Get("method").(string)

	verificationReq := siteverification.SiteVerificationWebResourceResource{
		Site: &siteverification.SiteVerificationWebResourceResourceSite{
			Identifier: d.Get("identifier").(string),
		},
	}
	log.Printf("[DEBUG] Site Verification insert request: %#v", verificationReq)

	resp, err := config.
		clientSiteVerification.
		WebResource.Insert(method, verificationReq).Do()
	if err != nil {
		return fmt.Errorf("Error requesting site verification: %#v", err)
	}
	log.Printf("[DEBUG] Site verification insert response: %#v", resp)

	err := d.Set("id", resp.Id)
	return err
}

func resourceSiteVerificationRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	readReq := siteverification.SiteVerificationWebResourceResource{
		Site: &siteverification.SiteVerificationWebResourceResourceSite{
			Identifier: d.Get("identifier").(string),
		},
	}
	// TODO logging
	resp, err := config.
		clientSiteVerification.
		WebResource.Get(d.Get("id").(string)).Do()

	// TODO logging
	return nil
}

func resourceSiteVerificationDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	deleteReq := siteverification.SiteVerificationWebResourceResource{
		Site: &siteverification.SiteVerificationWebResourceResourceSite{
			Identifier: d.Get("identifier").(string),
		},
	}
	// TODO logging

	resp, err := config.
		clientSiteVerification.
		WebResource.Delete(d.Get("id").(string)).Do()
	// TODO logging

	return nil
}
