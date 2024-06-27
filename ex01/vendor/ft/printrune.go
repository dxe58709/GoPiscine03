/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printrune.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: nsakanou <nsakanou@student.42tokyo.jp>     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/27 17:44:22 by nsakanou          #+#    #+#             */
/*   Updated: 2024/06/27 17:44:23 by nsakanou         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package ft

import (
	"errors"
	"os"
	"unicode/utf8"
)

// PrintRune prints a single rune (Unicode code point) and returns any error
// if the encoding or the writing fails.
func PrintRune(r rune) error {
	l := utf8.RuneLen(r)
	if l == -1 {
		return errors.New("The rune is not a valid value to encode in UTF-8")
	}
	p := make([]byte, l)
	utf8.EncodeRune(p, r)
	_, err := os.Stdout.Write(p)
	return err
}