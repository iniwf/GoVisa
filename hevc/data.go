/* The copyright in this software is being made available under the BSD
 * License, included below. This software may be subject to other third party
 * and contributor rights, including patent rights, and no such rights are
 * granted under this license.
 *
 * Copyright (c) 2012-2013, H265.net
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *  * Redistributions of source code must retain the above copyright notice,
 *    this list of conditions and the following disclaimer.
 *  * Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 *  * Neither the name of the H265.net nor the names of its contributors may
 *    be used to endorse or promote products derived from this software without
 *    specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS
 * BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF
 * THE POSSIBILITY OF SUCH DAMAGE.
 */
 
package hevc

import (
	"fmt"
    "io"
    "bufio"
    "errors"
    "strings"
    "strconv"
)


const (
    TCOEF = iota
    TPEL
    TPXL
)

type DATA struct {
	YUVReader	io.Reader
	Reader      *bufio.Reader
	//M_frame			   *Frame
    //M_tu               *TU
    M_data			   [][]int
    M_color				int
    M_tu_x				int
    M_tu_y				int
    M_width				int
    M_height			int
    M_type				int
}


func NewDATA(reader *bufio.Reader, tu_color, tu_x, tu_y, tu_width, tu_height, pixel_type int) *DATA{
	var ps DATA
	
	ps.Reader = reader
	ps.M_color = tu_color
	ps.M_tu_x = tu_x
	ps.M_tu_y = tu_y
	ps.M_width = tu_width
	ps.M_height = tu_height
	ps.M_type = pixel_type
	
	ps.M_data = make([][]int, tu_height)
	for j:= 0; j<tu_height; j=j+1 {
		ps.M_data[j] = make([]int, tu_width)
	}
	
    return &ps
}

func (ps *DATA) ParseYUV() (err error)  {
	var n int
	data := make([]byte, ps.M_width)
	
	err = nil
	for j:=0; j<ps.M_height; j=j+1{
		n, err = ps.YUVReader.Read(data)
		if n!=ps.M_width {
			err = errors.New("read incomplete frame data "+strconv.Itoa(j)+","+strconv.Itoa(n)+","+strconv.Itoa(ps.M_width)+"\n")
			break
		}
		
		for i:=0; i<ps.M_width; i=i+1{
			ps.M_data[j][i] = int(data[i])
		}
	}
	return err
}

func (ps *DATA) Parse() (line string, err error)  {
    if ps.Reader==nil{
		line = ""
		err = errors.New("nil Reader\n")
		return line, err
	}

	j := 0
    eof := false
    for !eof {
        line, err = ps.Reader.ReadString('\n')
        if err == io.EOF {
            err = nil
            eof = true
        }else if err != nil {
            return line, err
        }else if strings.HasPrefix(line, "=========") {
            return line, err
        }else if j>=ps.M_height {
        	line = ""
        	err =  errors.New("more data than expected")
        	return line, err
        }else{
        	var value	int64
            coef_str   := strings.Fields(strings.TrimSpace(line))
			
			for i:=0; i<ps.M_width; i=i+1{
            	value, err = strconv.ParseInt(coef_str[i], 16, 0)
            	if err != nil {
                	return line, err
            	}else{
            		ps.M_data[j][i] = int(value)
            	}
            }
            j = j+1
        }
    }
        
    return line, err
}


func (ps *DATA) ShowInfo(){
	var color	string
	if ps.M_color==0 {
		color = " Y"
	}else if ps.M_color==1 {
		color = "Cb"
	}else{
		color = "Cr"
	}
	 
    fmt.Printf("========= TU (%s,%2d,%2d) ===================================================\n", color, ps.M_tu_x, ps.M_tu_y)
    for j:=0; j<ps.M_height; j=j+1 {
    	for i:=0; i<ps.M_width; i=i+1{
    		if ps.M_type == TCOEF {
        		fmt.Printf("%04x ", ps.M_data[j][i])
        	}else if ps.M_type == TPEL {
        		fmt.Printf("%04x ", ps.M_data[j][i])
        	}else {//if ps.M_type == TPXL {
        		fmt.Printf("%02x ", ps.M_data[j][i])
        	}
        }
        fmt.Printf("\n")
    }
  //fmt.Printf("===========================================================================\n")
}
