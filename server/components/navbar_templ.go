// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.754
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Navbar() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<navbar class=\"navbar bg-base-100\"><div class=\"navbar-start\"><div class=\"dropdown\"><div tabindex=\"0\" role=\"button\" class=\"btn btn-ghost lg:hidden\"><svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-5 w-5\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M4 6h16M4 12h8m-8 6h16\"></path></svg></div><ul tabindex=\"0\" class=\"menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow\"><li><a>Item 1</a></li><li><a>Parent</a><ul class=\"p-2\"><li><a>Submenu 1</a></li><li><a>Submenu 2</a></li></ul></li><li><a>Item 3</a></li></ul></div><a class=\"btn btn-ghost text-xl\">Zelvy</a></div><div class=\"navbar-center hidden lg:flex\"><ul class=\"menu menu-horizontal px-1\"><li><a>Item 1</a></li><li><details><summary>Parent</summary><ul class=\"p-2\"><li><a>Submenu 1</a></li><li><a>Submenu 2</a></li></ul></details></li><li><a>Item 3</a></li></ul></div></navbar>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
