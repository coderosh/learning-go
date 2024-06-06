// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Nav(userId string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav class=\"w-full h-16 sticky top-0 py-2 border-b bg-white\"><div class=\"flex px-4 justify-between  items-center\"><a class=\"text-lg font-bold text-violet-500\" href=\"/\">GSHORT</a><div></div><div class=\"flex gap-3\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if userId != "" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"inline-flex items-center justify-center px-4 py-2 text-sm font-medium tracking-wide text-white transition-colors duration-200 rounded-md bg-neutral-950 hover:bg-neutral-900 focus:ring-2 focus:ring-offset-2 focus:ring-neutral-900 focus:shadow-outline focus:outline-none\" href=\"/dashboard\">Dashboard</a><form action=\"/logout\" method=\"post\"><button type=\"submit\" class=\"nline-flex items-center justify-center px-4 py-2 text-sm font-medium tracking-wide transition-colors duration-100 bg-white border-2 rounded-md text-neutral-900 hover:text-white border-neutral-900 hover:bg-neutral-900\">Logout</button></form>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<a class=\"inline-flex items-center justify-center px-4 py-2 text-sm font-medium tracking-wide text-white transition-colors duration-200 rounded-md bg-neutral-950 hover:bg-neutral-900 focus:ring-2 focus:ring-offset-2 focus:ring-neutral-900 focus:shadow-outline focus:outline-none\" href=\"/login\">Login</a> <a class=\"inline-flex items-center justify-center px-4 py-2 text-sm font-medium tracking-wide transition-colors duration-100 bg-white border-2 rounded-md text-neutral-900 hover:text-white border-neutral-900 hover:bg-neutral-900\" href=\"/signup\">Signup</a>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
