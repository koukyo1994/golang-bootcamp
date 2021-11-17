package main_test

import (
	"strings"

	main "bootcamp/ch05/ex5.7"

	"golang.org/x/net/html"
)

func ExampleForEachNode() {
	doc, _ := html.Parse(strings.NewReader(`
	<html lang="en">
		<head>
			<title>image</title>
		</head>
		<body>
			<!-- コメント -->
			<h1>リンクのサンプル</h1>
			<p><a href="https://gopl.io">Gopl.io</a></p>
			<h1>画像のサンプル</h1>
			<img src='image.png'>
			<h2>テーブルのサンプル</h2>
			<table>
				<tr style='text-align: left'>
					<th>Item</th>
					<th>Price</th>
				</tr>
				<tr>
					<td>aaa</td>
					<td>$10</td>
				</tr>
			</table>
		</body>
	</html>`))
	main.ForEachNode(doc, main.StartElement, main.EndElement)
	// Output:
	// <html lang='en'>
	//   <head>
	//     <title>
	//       image
	//     </title>
	//   </head>
	//   <body>
	//     <!-- コメント -->
	//     <h1>
	//       リンクのサンプル
	//     </h1>
	//     <p>
	//       <a href='https://gopl.io'>
	//         Gopl.io
	//       </a>
	//     </p>
	//     <h1>
	//       画像のサンプル
	//     </h1>
	//     <img src='image.png'/>
	//     <h2>
	//       テーブルのサンプル
	//     </h2>
	//     <table>
	//       <tbody>
	//         <tr style='text-align: left'>
	//           <th>
	//             Item
	//           </th>
	//           <th>
	//             Price
	//           </th>
	//         </tr>
	//         <tr>
	//           <td>
	//             aaa
	//           </td>
	//           <td>
	//             $10
	//           </td>
	//         </tr>
	//       </tbody>
	//     </table>
	//   </body>
	// </html>
}
