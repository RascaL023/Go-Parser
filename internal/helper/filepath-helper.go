package helper

import (
	"os"
	"parsertry/internal/state"
	"strings"
)

func parseKeyword(keyword string, state *state.State) string {
	res := strings.ToUpper(keyword);

	switch res {
		case "WAYBAR" : res = state.Waybar;
		case "THEME"  : res = state.Theme.Name;
		default: 
			if res = os.Getenv(res); res == "" {
				res = "$" + keyword;
			}
	}
	
	return res;
}

// Faster - minimum alloc
func ExpandPath(path string, state *state.State) string {
    // small buffer di stack (tidak alloc heap kalo < 64 byte)
    buf := make([]byte, 0, len(path)+32);

    for i := 0; i < len(path); {
        if path[i] != '$' {
            buf = append(buf, path[i]);
            i++;
            continue;
        }

        start := i + 1;
        end := start;

        for end < len(path) {
            ch := path[end];
            if ch == '/' || ch == '|' || ch == '\\' || ch == '.' {
                break;
            }
            end++;
        }

				parsed := parseKeyword(path[start:end], state);

        buf = append(buf, parsed...);

        i = end;
    }

    return string(buf);
}


// Faster
// func ExpandPath(path string, state *state.State) string {
//     var builder strings.Builder
//     builder.Grow(len(path) + 32)
//
//     for i := 0; i < len(path); {
//         if path[i] != '$' {
//             builder.WriteByte(path[i]);
//             i++;
//             continue;
//         }
//         start := i + 1;
//         end := start;
//
//         for end < len(path) {
//             ch := path[end];
//             if ch == '/' || ch == '|' || ch == '\\' || ch == '.' {
//                 break;
//             }
//             end++;
//         }
// 				parsed := parseKeyword(path[start:end], state);
//         builder.WriteString(parsed);
//
//         i = end;
//     }
//
//     return builder.String();
// }


// Mine
// func ExpandPath(path string, state *state.State) string {
// 	for i := 0; i < len(path); i++ {
// 		if path[i] == '$' {
// 			start  := i + 1;
// 			end 	 := start;
// 			parsed := "";
//
// 			for end < len(path) {
// 				ch := path[end];
// 				if ch == '/' || ch == '|' || ch == '\\' || ch == '.' {
// 					parsed = parseKeyword(path[start:end], state);
// 					path = path[:start - 1] + parsed + path[end:];
//
// 					i = end - 1;
// 					break;
// 				}
//
// 				end++;
// 			}
// 		}
// 	}
//
// 	return path;
// }
