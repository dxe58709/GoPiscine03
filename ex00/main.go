/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/29 16:09:02 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/29 16:18:05 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"ft"
	"piscine"
)

func main() {
	ft.PrintRune(piscine.FirstRune("Hello!"))
	ft.PrintRune(piscine.FirstRune("Salut!"))
	ft.PrintRune(piscine.FirstRune("Ola!"))
	ft.PrintRune('\n')
}
