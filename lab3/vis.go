package main

// // import (
// // 	"fmt"
// // )

// // // func printBohr() {
// // // 	fmt.Println("\n=== Текущее состояние бора ===")
// // // 	fmt.Printf("%-5s | %-20s | %-10s | %-10s | %-10s | %-10s\n",
// // // 		"Node", "Transitions (A,C,G,T,N)", "SuffLink", "SuffFlink", "Pattern", "Terminal")

// // // 	for i, node := range bohr {
// // // 		transitions := ""
// // // 		for ch, next := range node.next_vrtx {
// // // 			if next != -1 {
// // // 				transitions += fmt.Sprintf("%c->%d ", "ACGTN"[ch], next)
// // // 			}
// // // 		}

// // // 		patternInfo := ""
// // // 		if node.flag {
// // // 			patternInfo = pattern[node.pat_num]
// // // 		}

// // // 		fmt.Printf("%-5d | %-20s | %-10d | %-10d | %-10s | %-10t\n",
// // // 			i, transitions, node.suff_link, node.suff_flink, patternInfo, node.flag)
// // // 	}
// // // 	fmt.Println("=============================")
// // // }

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"text/template"
// )

// var tmpl = `
// <!DOCTYPE html>
// <html>
// <head>
//     <title>Визуализация бора</title>
//     <script src="https://d3js.org/d3.v7.min.js"></script>
//     <style>
//         .node circle {
//             fill: #fff;
//             stroke: steelblue;
//             stroke-width: 2px;
//         }
//         .node text { font: 12px sans-serif; }
//         .link {
//             fill: none;
//             stroke: #ccc;
//             stroke-width: 1.5px;
//         }
//         .suffix-link {
//             fill: none;
//             stroke: red;
//             stroke-width: 1px;
//             stroke-dasharray: 5,5;
//         }
//     </style>
// </head>
// <body>
//     <div id="tree"></div>
//     <script>
//         const data = {{.GraphData}};

//         const width = 800, height = 600;
//         const svg = d3.select("#tree").append("svg")
//             .attr("width", width)
//             .attr("height", height);

//         const g = svg.append("g").attr("transform", "translate(40,0)");

//         const treeLayout = d3.tree().size([height-100, width-200]);
//         const root = d3.hierarchy(data);
//         treeLayout(root);

//         // Отрисовка связей
//         g.selectAll(".link")
//             .data(root.links())
//             .enter().append("path")
//             .attr("class", "link")
//             .attr("d", d3.linkVertical()
//                 .x(d => d.x)
//                 .y(d => d.y));

//         // Отрисовка узлов
//         const node = g.selectAll(".node")
//             .data(root.descendants())
//             .enter().append("g")
//             .attr("class", "node")
//             .attr("transform", d => ` + "`translate(${d.x},${d.y})`" + `);

//         node.append("circle")
//             .attr("r", 10)
//             .attr("fill", d => d.data.terminal ? "#aaffaa" : "#fff");

//         node.append("text")
//             .attr("dy", ".35em")
//             .attr("y", d => d.children ? -20 : 20)
//             .style("text-anchor", "middle")
//             .text(d => d.data.name);
//     </script>
// </body>
// </html>
// `

// func visualizeHandler(w http.ResponseWriter, r *http.Request) {
// 	data := prepareGraphData()
// 	t := template.Must(template.New("vis").Parse(tmpl))
// 	t.Execute(w, map[string]interface{}{"GraphData": data})
// }

// // Подготавливает данные для визуализации в D3.js
// func prepareGraphData() map[string]interface{} {
// 	nodes := make([]map[string]interface{}, len(bohr))
// 	links := make([]map[string]interface{}, 0)

// 	// Создаем узлы
// 	for i, node := range bohr {
// 		nodeData := map[string]interface{}{
// 			"id":    i,
// 			"name":  fmt.Sprintf("%d", i),
// 			"label": fmt.Sprintf("%d", i),
// 		}

// 		// Добавляем информацию о терминальных узлах
// 		if node.flag {
// 			nodeData["terminal"] = true
// 			nodeData["pattern"] = pattern[node.pat_num]
// 			nodeData["label"] = fmt.Sprintf("%d\n%s", i, pattern[node.pat_num])
// 		}

// 		// Добавляем информацию о корневом узле
// 		if i == 0 {
// 			nodeData["root"] = true
// 		}

// 		nodes[i] = nodeData
// 	}

// 	// Создаем связи (обычные переходы)
// 	for i, node := range bohr {
// 		for ch, next := range node.next_vrtx {
// 			if next != -1 {
// 				links = append(links, map[string]interface{}{
// 					"source": i,
// 					"target": next,
// 					"label":  string("ACGTN"[ch]),
// 					"type":   "transition",
// 				})
// 			}
// 		}
// 	}

// 	// Создаем суффиксные ссылки
// 	for i := range bohr {
// 		if suff := get_suff_link(i); suff != -1 && suff != i && suff != 0 {
// 			links = append(links, map[string]interface{}{
// 				"source": i,
// 				"target": suff,
// 				"type":   "suffix",
// 			})
// 		}
// 	}

// 	return map[string]interface{}{
// 		"nodes": nodes,
// 		"links": links,
// 	}
// }

// // func main() {
// // 	// Инициализация и построение бора
// // 	bohr_ini()
// // 	add_string_to_bohr("AC")
// // 	add_string_to_bohr("T")

// // 	// Запуск веб-сервера
// // 	http.HandleFunc("/", visualizeHandler)
// // 	fmt.Println("Сервер запущен на http://localhost:8080")
// // 	http.ListenAndServe(":8080", nil)
// // }
