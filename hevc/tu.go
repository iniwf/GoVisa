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
    //"fmt"
    //"errors"
    "bufio"
    "container/list"
)


type TU struct {
    ParameterSet
    CoefList		   *list.List
    ResiList		   *list.List
    PredList		   *list.List
    RecoList		   *list.List
    FinalList		   *list.List
    M_cu               *CU
    Tu_color			int
    Tu_x				int
    Tu_y				int
    Tu_width			int
	Tu_height			int
}


func NewTU(reader *bufio.Reader) *TU{
	var tu TU
	
	tu.ParameterSet.Reader = reader
	tu.ParameterSet.FieldList = list.New()
	tu.CoefList = list.New()
	tu.ResiList = list.New()
	tu.PredList = list.New()
	tu.RecoList = list.New()
	tu.FinalList= list.New()
	
    return &tu
}

func (ps *TU) Parse() (line string, err error)  {
    line, err = ps.ParameterSet.Parse()

	ps.Tu_color, err = ps.ParameterSet.GetValue("tu_color") 
    if err != nil {
    	return line, err
    }
    
	ps.Tu_x, err = ps.ParameterSet.GetValue("tu_x") 
    if err != nil {
    	return line, err
    }
    
    ps.Tu_y, err = ps.ParameterSet.GetValue("tu_y") 
    if err != nil {
    	return line, err
    }
    
	ps.Tu_width, err = ps.ParameterSet.GetValue("tu_width") 
    if err != nil {
    	return line, err
    }
    
    ps.Tu_height, err = ps.ParameterSet.GetValue("tu_height") 
    
    return line, err
}


