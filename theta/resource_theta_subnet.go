package theta

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"reflect"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSubnet() *schema.Resource {
	return &schema.Resource{
		// CRUD (CREATE, READ, UPDATE, DELETE) operations.
		//
		// When terraform reads the state file, if a resource doesn't exist, then a
		// CREATE operation will be started (otherwise an UPDATE operation).
		//
		Create: resourceSubnetCreate,
		Read:   resourceSubnetRead,
		Update: resourceSubnetUpdate,
		Delete: resourceSubnetDelete,

		// Resource Schema
		//
		// NOTE:
		// You must specify either 'optional', 'required', or 'computed' and the
		// value needs to be set to boolean 'true'.
		//
		// Reference:
		// https://www.terraform.io/docs/extend/schemas/schema-types.html
		Schema: map[string]*schema.Schema{
			"last_updated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fn_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"not_computed_optional": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"handler": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"filename": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"function_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"runtime": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"rpc": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tfuel_token_bank_addr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tfuel_tnt20_bank_addr": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"initial_fee": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"cross_chain_transfer_fee_in_tfuel": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"key": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			// This attribute is 'required' meaning the consumer of this provider
			// will need to define the values expected when writing their terraform
			// HCL code.
			"foo": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bar": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"number": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"version": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			/*
				foo example:

				resource "mock_example" "testing" {
					foo {
						bar {
							number = 1
						}
					}
				  foo {
						bar {
							number = 2
						}
					}
					foo {
						bar {
							number = 3
						}
					}
				}

				OR

				resource "mock_example" "testing" {
					dynamic "foo" {
						for_each = [{ number = 1 }, { number = 2 }, { number = 3 }]
						content {
							bar {
								number = foo.value.number
							}
						}
					}
				}
			*/
			"baz": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"qux": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			/*
				baz example:

				resource "mock_example" "testing" {
					baz {
						qux = "x"
					}
					baz {
						qux = "y"
					}
					baz {
						qux = "z"
					}
				}

				OR

				resource "mock_example" "testing" {
					dynamic "baz" {
						for_each = [{ qux = "x" }, { qux = "y" }, { qux = "z" }]
						content {
							qux = baz.value.qux
						}
					}
				}
			*/

			"some_list": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func exec_cmd(acwd string, acmd []string) ([]byte, error) {
	fun := reflect.ValueOf(exec.Command)
	var args []reflect.Value
    for _, x := range acmd {
        args = append(args, reflect.ValueOf(x))
    }
    result := fun.Call(args)
	cmd_ := result[0].Interface().(*exec.Cmd)
	cmd_.Dir = acwd
	stdout_, stderr_ := cmd_.Output()
	log.Printf("\n\n>>> stdout %+v\n\n", stdout_)
	log.Printf("\n\n>>> stderr: %+v\n\n", stderr_)

	return stdout_, stderr_
}

func subnet_create(d *schema.ResourceData, m any) error {
	fmt.Print("\n\n--- CREATE THETA SUBNET ---\n\n")
	log.Printf("\n\n>>> schema.ResourceData: %+v\n\n", d)
	log.Printf("\n\n>>> meta data: %+v\n\n", m)

	// return nil

	// theta functions create --name my-first-function
	// theta functions deploy \
    // --name my-first-function \
    // --path ~/some/path/my-first-function.js

	const METACHAIN_GUIDE_ROOT = "~/metachain_playground/theta-metachain-guide"
	const SDK_JS_DIR = METACHAIN_GUIDE_ROOT + "sdk/js"
	const WS_DIR = "~/metachain_playground/privatenet/workspace"
	
	// id := d.Get("id").(string);
	// name := d.Get("name").(string);
	// rpc := d.Get("rpc").(string);
	// tfuel_token_bank_addr := d.Get("tfuel_token_bank_addr").(string);
	// tfuel_tnt20_bank_addr := d.Get("tfuel_tnt20_bank_addr").(string);
	// initial_fee := d.Get("initial_fee").(int);
	key := d.Get("key").(string);


	pw := "qwertyuiop"

	validator := "0x2E833968E5bB786Ae419c4d13189fB081Cc43bab"

	var cmd_ []string
	// var stdout_ []byte
	// var stderr_ error -
	
	// 2. Deploy the Governance Token for the Subchain
	//cd $METACHAIN_GUIDE_ROOT cd sdk/js
	// node deployGovToken.js privatenet "Subchain 360777 Gov" GOV360777 0x2E833968E5bB786Ae419c4d13189fB081Cc43bab 0x2E833968E5bB786Ae419c4d13189fB081Cc43bab ~/.thetacli/keys/encrypted/2E833968E5bB786Ae419c4d13189fB081Cc43bab qwertyuiop

	cmd_ = []string {"node", "deployGovToken.js", "privatenet", "Subchain 360777 Gov", "GOV360777", validator, validator, key, pw}
	stdout_, stderr_ := exec_cmd(SDK_JS_DIR, cmd_)


	// 3. Register a New Subchain
	// cd $METACHAIN_GUIDE_ROOT cd sdk/js
	// node mintMockWrappedTheta.js privatenet 0x2E833968E5bB786Ae419c4d13189fB081Cc43bab 50000000000000000000000 ~/.thetacli/keys/encrypted/2E833968E5bB786Ae419c4d13189fB081Cc43bab qwertyuiop

	balance := "50000000000000000000000"

	cmd_ = []string {"node", "mintMockWrappedTheta.js", "privatenet", "GOV360777", validator, balance, key, pw}
	stdout_, stderr_ = exec_cmd(SDK_JS_DIR, cmd_)

	balance2 := "100000000000000000000000"

	// 4. Stake to a New Validator
	// cd $METACHAIN_GUIDE_ROOT cd sdk/js
	// node depositStake.js privatenet 100000000000000000000000 0x2E833968E5bB786Ae419c4d13189fB081Cc43bab ~/.thetacli/keys/encrypted/2E833968E5bB786Ae419c4d13189fB081Cc43bab qwertyuiop

	cmd_ = []string {"node", "depositStake.js", "privatenet", balance2, validator, key, pw}
	stdout_, stderr_ = exec_cmd(SDK_JS_DIR, cmd_)

	// 5. Run the Subchain validator and the ETH RPC Adapter
	// cd ~/metachain_playground/privatenet/workspace
	// theta-eth-rpc-adaptor start --config=../subchain/ethrpc
	// cd ~/metachain_playground/privatenet/workspace
	// thetasubchain start --config=../subchain/validator --password=qwertyuiop

	cmd_ = []string {"theta-eth-rpc-adaptor", "start", "--config="+"../subchain/ethrpc"}
	stdout_, stderr_ = exec_cmd(WS_DIR, cmd_)

	cmd_ = []string {"thetasubchain", "start", "--config="+"../subchain/validator", "--password="+pw}
	stdout_, stderr_ = exec_cmd(WS_DIR, cmd_)

	log.Printf("\n\n>>> stdout %+v\n\n", stdout_)
	log.Printf("\n\n>>> stderr: %+v\n\n", stderr_)

	// var outs string = string(stdout1[:]);
	// outs += "\n______\n";
	// outs += string(stdout2[:]);
	// outs += "\n______\n";
	
	// lines := strings.Split(string(stdout2[:]), "\n");
	// outs := lines[len(lines)-2];


	// // var fn_id string = strings.Join(lines, "\n__ ");
	// var fn_id string = outs;
	// d.Set("fn_id", fn_id);

	return nil;
}

// The CREATE operation typically will make API calls to create resources. It
// won't set anything in the terraform state, with the exception of setting a
// unique ID that will be used by all the other functions to access the
// resource data from state.
func resourceSubnetCreate(d *schema.ResourceData, m any) error {
	log.Print("\n\n--- CREATE ---\n\n")
	log.Printf("\n\n>>> schema.ResourceData: %+v\n\n", d)
	log.Printf("\n\n>>> meta data: %+v\n\n", m)

	// If this were a real provider, then we'd have an API client that would be
	// creating, reading, updating, deleting (i.e. CRUD) data.
	//
	// c := m.(*some.APIClient)

	foo := d.Get("foo").([]any)
	log.Printf(">>> foo: %+v (%T)\n", foo, foo)

	// We could do things with foo, e.g. loop over its elements and build up a
	// data structure to be used as input to an API client to create.
	//
	// Remember that "foo" was defined in the schema as "required" meaning the
	// consumer of this provider would need to provide the values associated with
	// the foo schema.

	// Let's pretend we made an API call and in the response we got back there
	// was an ID we could use as a unique key in our terraform state file so it
	// was able to track the resource.
	//
	// NOTE:
	// The mere existence of the ID and lack of error means terraform will
	// presume the CREATE operation was successful and store the provided "foo"
	// data the user provided into the local state file.
	d.SetId("123")

	// We do a READ operation to be sure we get the latest state stored locally.
	//
	// NOTE:
	// In this example, instead of doing a READ and then returning nil at the end
	// of the create function, we instead return the result of the READ. Meaning
	// if the READ fails, then we'll fail the CREATE. It's up to you whether
	// that's something you want to do as a READ could fail due to a network
	// issue and not necessarily mean there was an error with the CREATE
	// operation (which itself would have caused an error earlier and failed the
	// CREATE any way). This way we're ensuring the local state is up-to-date and
	// doesn't need a refresh.
	return subnet_create(d, m)
}

// The READ operation must handle three things: calling out to the API to get
// the latest version of our resource, and to flatten the returned API response
// (or at least extract the values we need) and marshal the data into a format
// that terraform understands (i.e. map[string]any). Lastly it will set
// the latest data into terraform's state file so terraform can identify if
// there are any differences between what the user has defined and what
// actually exists in reality.
func resourceSubnetRead(d *schema.ResourceData, m any) error {
	log.Print("\n\n--- READ ---\n\n")
	log.Printf("\n\n>>> schema.ResourceData: %+v\n\n", d)
	log.Printf("\n\n>>> meta data: %+v\n\n", m)

	// We get the ID we set into terraform state after we had initially created
	// the resource.
	resourceID := d.Id()
	log.Println(">>> resourceID:", resourceID)

	// For the purposes of our 'mock' provider, we don't have an API to get data
	// from and so I'm going to hardcode data to be set into terraform state. To
	// do that I have to first get the data out from terraform state.
	foo := d.Get("foo")
	log.Printf("\n\n>>> foo: %+v (%T)\n\n", foo, foo)

	// I want to set a computed value for the nested 'version' attribute, but to
	// do that I have to iterate over each parent structure until I reach the
	// relevant level of the data structure where I can then set a value on the
	// 'version' attribute.

	// In order to loop over foo, we need to cast it to the appropriate type
	for _, f := range foo.([]any) {
		f := f.(map[string]any)
		log.Printf("\n\n>>> f: %+v (%T)\n\n", f, f)

		// In order to loop over the nested bar, we need to cast it to the appropriate type
		for _, b := range f["bar"].([]any) {
			log.Printf("\n\n>>> b: %+v (%T)\n\n", b, b)

			// In order to update the 'version' field inside of 'bar', we need to
			// cast bar to the appropriate type first
			b := b.(map[string]any)
			b["version"] = uuid.New().String()
		}
	}

	log.Printf("\n\n>>> foo (AFTER UPDATE): %+v (%T)\n\n", foo, foo)

	// Now we can set the updated data structure into local state.
	if err := d.Set("foo", foo); err != nil {
		return err
	}

	// I also make sure to update the computed 'last_update' attribute every time
	// we update the terraform state.
	d.Set("last_updated", time.Now().Format(time.RFC850))

	return nil
}

func resourceSubnetUpdate(d *schema.ResourceData, m any) error {
	log.Print("\n\n--- UPDATE ---\n\n")
	log.Printf("\n\n>>> schema.ResourceData: %+v\n\n", d)
	log.Printf("\n\n>>> meta data: %+v\n\n", m)

	// We get the ID we set into terraform state after we had initially created
	// the resource.
	resourceID := d.Id()
	log.Println("resourceID:", resourceID)

	if d.HasChange("foo") {
		foo := d.Get("foo").([]any)
		log.Printf(">>> foo: %+v\n", foo)

		// Imagine we made an API call to update the given resource.
		//
		// We'd do this by iterating over the foo we pulled out of our terraform
		// state and coercing them into a type of map[string]any
		//
		// e.g.
		//
		// for _, f := range foo {
		// 	i := f.(map[string]any)
		//
		// 	t := i["bar"].([]any)[0]
		// 	bar := t.(map[string]any)
		//
		//  ...constructing data structure to pass to API...
		//
		//  We might assign values to the data structure like:
		//
		//  bar["id"].(int)).
		// }

		// TODO: update "version" to be 2

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	// Again, we do a READ operation to be sure we get the latest state stored locally.
	//
	// return resourceRead(d, m)
	return theta_edge_create(d, m);

}

func resourceSubnetDelete(d *schema.ResourceData, m any) error {
	log.Print("\n\n--- DELETE ---\n\n")
	log.Printf("\n\n>>> schema.ResourceData: %+v\n\n", d)
	log.Printf("\n\n>>> meta data: %+v\n\n", m)

	// We get the ID we set into terraform state after we had initially created
	// the resource.
	resourceID := d.Id()
	log.Println(">>> resourceID:", resourceID)

	// Imagine we use resourceID to issue a DELETE API call.

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return nil
}
