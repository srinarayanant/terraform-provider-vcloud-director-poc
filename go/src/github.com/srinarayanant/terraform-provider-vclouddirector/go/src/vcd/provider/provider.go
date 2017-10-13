package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	//"log"
	//"os"
)

// .. Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"user": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCD_USER", nil),
				Description: "The user name for vcd API operations.",
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCD_PASSWORD", nil),
				Description: "The user password for vcd API operations.",
			},

			"org": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCD_ORG", nil),
				Description: "The vcd org for API operations",
			},

			"ip": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCD_IP", nil),
				Description: "The vcd IP for vcd API operations.",
			},

			"vdc": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCD_VDC", ""),
				Description: "The name of the VDC to run operations on",
			},

			"maxRetryTimeout": &schema.Schema{
				Type:       schema.TypeInt,
				Optional:   true,
				Deprecated: "Deprecated. Use max_retry_timeout instead.",
			},

			"max_retry_timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCD_MAX_RETRY_TIMEOUT", 60),
				Description: "Max num seconds to wait for successful response when operating on resources within vCloud (defaults to 60)",
			},

			"allow_unverified_ssl": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VCD_ALLOW_UNVERIFIED_SSL", false),
				Description: "If set, VCDClient will permit unverifiable SSL certificates.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{

			"vclouddirector_catalog": resourceCatalog(),
			//			"vcd_network":         resourceVcdNetwork(),
			//			"vcd_vapp":            resourceVcdVApp(),
			//			"vcd_firewall_rules":  resourceVcdFirewallRules(),
			//			"vcd_dnat":            resourceVcdDNAT(),
			//			"vcd_snat":            resourceVcdSNAT(),
			//			"vcd_edgegateway_vpn": resourceVcdEdgeGatewayVpn(),
			//			"vcd_vapp_vm":         resourceVcdVAppVm(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	maxRetryTimeout := d.Get("max_retry_timeout").(int)

	//	log.SetOutput(os.Stdout)

	// TODO: Deprecated, remove in next major release
	if v, ok := d.GetOk("maxRetryTimeout"); ok {
		maxRetryTimeout = v.(int)
	}

	config := Config{
		User:            d.Get("user").(string),
		Password:        d.Get("password").(string),
		Org:             d.Get("org").(string),
		Ip:              d.Get("ip").(string),
		VDC:             d.Get("vdc").(string),
		MaxRetryTimeout: maxRetryTimeout,
		InsecureFlag:    d.Get("allow_unverified_ssl").(bool),
	}

	vcdclient, err := config.CreateClient()
	return vcdclient, err
}
