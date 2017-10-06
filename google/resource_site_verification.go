package google

// import ()

// func resourceSiteVerification() *schema.Resource {
// 	return &schema.Resource{
// 		Create: resourceSiteVerificationCreate,
// 		Read:   resourceSiteVerificationRead,
// 		Delete: resourceSiteVerificationDelete,

// 		Schema: map[string]*schema.Schema{
// 			"identifier": &schema.Schema{
// 				Type:        schema.TypeString,
// 				Required:    true,
// 				Description: "Domain name to verfiy",
// 				ForceNew:    true,
// 			},
// 			"method": &schema.Schema{
// 				Type:         schema.TypeString,
// 				Required:     true,
// 				ForceNew:     true,
// 				ValidateFunc: siteVerificationMethodValid,
// 			},
// 			"id": &schema.Schema{
// 				Type:     schema.TypeString,
// 				Computed: true,
// 				ForceNew: true,
// 			},
// 		},
// 	}
// }

// func resourceSiteVerificationCreate(d *schema.ResourceData, meta interface{}) error {
// 	config := meta.(*Config)

// 	return nil
// }

// func resourceSiteVerificationRead(d *schema.ResourceData, meta interface{}) error {
// 	config := meta.(*Config)

// 	return nil
// }

// func resourceSiteVerificationDelete(d *schema.ResourceData, meta interface{}) error {
// 	config := meta.(*Config)

// 	return nil
// }
