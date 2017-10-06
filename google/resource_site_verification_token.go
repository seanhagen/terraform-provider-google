package google

import (
	"log"

	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"google.golang.org/api/siteverification/v1"
	"strings"
)

func resourceSiteVerificationToken() *schema.Resource {
	return &schema.Resource{
		Create: resourceSiteVerificationTokenCreate,
		Read:   resourceSiteVerificationTokenRead,
		Delete: resourceSiteVerificationTokenDelete,
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
			"cname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func siteVerificationMethodValid(value interface{}, field string) (values []string, es []error) {
	test := value.(string)
	if field == "method" && test != "DNS_TXT" && test != "DNS_CNAME" {
		es = append(es, fmt.Errorf("Valid 'method' are as follows: 'DNS_TXT' or 'DNS_CNAME'"))
	}
	return
}

func resourceSiteVerificationTokenCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	method := d.Get("method").(string)

	verification := siteverification.SiteVerificationWebResourceGettokenRequest{
		Site: &siteverification.SiteVerificationWebResourceGettokenRequestSite{
			Identifier: d.Get("identifier").(string),
			Type:       "INET_DOMAIN",
		},
		VerificationMethod: method,
	}

	log.Printf("[DEBUG] Site Verification create token request: %#v", verification)
	resp, err := config.clientSiteVerification.WebResource.GetToken(&verification).Do()
	if err != nil {
		return fmt.Errorf("Error creating Site Verification token: %s", err)
	}

	log.Printf("[DEBUG] Site Verification token response: %#v", resp)
	token := ""
	cname := ""
	if method == "DNS_CNAME" {
		parts := strings.Split(resp.Token, " ")
		cname = parts[0] + "." + d.Get("identifier").(string)
		token = parts[1]
	} else {
		cname = d.Get("identifier").(string)
		token = resp.Token
	}
	err = d.Set("cname", cname)
	if err != nil {
		return err
	}

	err = d.Set("token", token)

	d.SetId(cname)
	return err
}

func resourceSiteVerificationTokenRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceSiteVerificationTokenDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
