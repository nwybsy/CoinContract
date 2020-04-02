//Package CoinContract comment
// This file war generated by tars2go 1.1
// Generated from IndexPriceObj.tars
package CoinContract

import (
	"context"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
)

//IndexPriceObj struct
type IndexPriceObj struct {
	s m.Servant
}

//Add is the proxy function for the method defined in the tars file, with the context
func (_obj *IndexPriceObj) Add(A int32, B int32, C *int32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(A, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32(B, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "Add", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_int32(&(*C), 3, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//AddWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *IndexPriceObj) AddWithContext(ctx context.Context, A int32, B int32, C *int32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(A, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32(B, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "Add", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_int32(&(*C), 3, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//Sub is the proxy function for the method defined in the tars file, with the context
func (_obj *IndexPriceObj) Sub(A int32, B int32, C *int32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(A, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32(B, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "Sub", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_int32(&(*C), 3, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SubWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *IndexPriceObj) SubWithContext(ctx context.Context, A int32, B int32, C *int32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(A, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32(B, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "Sub", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_int32(&(*C), 3, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//GetIndexPrice is the proxy function for the method defined in the tars file, with the context
func (_obj *IndexPriceObj) GetIndexPrice(Symbol string, Value *string, _opt ...map[string]string) (ret string, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Symbol, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "GetIndexPrice", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_string(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*Value), 2, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//GetIndexPriceWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *IndexPriceObj) GetIndexPriceWithContext(ctx context.Context, Symbol string, Value *string, _opt ...map[string]string) (ret string, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Symbol, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "GetIndexPrice", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_string(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*Value), 2, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SetServant sets servant for the service.
func (_obj *IndexPriceObj) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *IndexPriceObj) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}
func (_obj *IndexPriceObj) setMap(l int, res *requestf.ResponsePacket, ctx map[string]string, sts map[string]string) {
	if l == 1 {
		for k, _ := range ctx {
			delete(ctx, k)
		}
		for k, v := range res.Context {
			ctx[k] = v
		}
	} else if l == 2 {
		for k, _ := range ctx {
			delete(ctx, k)
		}
		for k, v := range res.Context {
			ctx[k] = v
		}
		for k, _ := range sts {
			delete(sts, k)
		}
		for k, v := range res.Status {
			sts[k] = v
		}
	}
}

//AddServant adds servant  for the service.
func (_obj *IndexPriceObj) AddServant(imp _impIndexPriceObj, obj string) {
	tars.AddServant(_obj, imp, obj)
}

//AddServant adds servant  for the service with context.
func (_obj *IndexPriceObj) AddServantWithContext(imp _impIndexPriceObjWithContext, obj string) {
	tars.AddServantWithContext(_obj, imp, obj)
}

type _impIndexPriceObj interface {
	Add(A int32, B int32, C *int32) (ret int32, err error)
	Sub(A int32, B int32, C *int32) (ret int32, err error)
	GetIndexPrice(Symbol string, Value *string) (ret string, err error)
}
type _impIndexPriceObjWithContext interface {
	Add(ctx context.Context, A int32, B int32, C *int32) (ret int32, err error)
	Sub(ctx context.Context, A int32, B int32, C *int32) (ret int32, err error)
	GetIndexPrice(ctx context.Context, Symbol string, Value *string) (ret string, err error)
}

func Add(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var A int32
	err = _is.Read_int32(&A, 1, true)
	if err != nil {
		return err
	}
	var B int32
	err = _is.Read_int32(&B, 2, true)
	if err != nil {
		return err
	}
	var C int32
	if withContext == false {
		_imp := _val.(_impIndexPriceObj)
		ret, err := _imp.Add(A, B, &C)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impIndexPriceObjWithContext)
		ret, err := _imp.Add(ctx, A, B, &C)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = _os.Write_int32(C, 3)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func Sub(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var A int32
	err = _is.Read_int32(&A, 1, true)
	if err != nil {
		return err
	}
	var B int32
	err = _is.Read_int32(&B, 2, true)
	if err != nil {
		return err
	}
	var C int32
	if withContext == false {
		_imp := _val.(_impIndexPriceObj)
		ret, err := _imp.Sub(A, B, &C)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impIndexPriceObjWithContext)
		ret, err := _imp.Sub(ctx, A, B, &C)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = _os.Write_int32(C, 3)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func GetIndexPrice(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var Symbol string
	err = _is.Read_string(&Symbol, 1, true)
	if err != nil {
		return err
	}
	var Value string
	if withContext == false {
		_imp := _val.(_impIndexPriceObj)
		ret, err := _imp.GetIndexPrice(Symbol, &Value)
		if err != nil {
			return err
		}

		err = _os.Write_string(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impIndexPriceObjWithContext)
		ret, err := _imp.GetIndexPrice(ctx, Symbol, &Value)
		if err != nil {
			return err
		}

		err = _os.Write_string(ret, 0)
		if err != nil {
			return err
		}
	}

	err = _os.Write_string(Value, 2)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//Dispatch is used to call the server side implemnet for the method defined in the tars file. withContext shows using context or not.
func (_obj *IndexPriceObj) Dispatch(ctx context.Context, _val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
	_is := codec.NewReader(tools.Int8ToByte(req.SBuffer))
	_os := codec.NewBuffer()
	switch req.SFuncName {
	case "Add":
		err := Add(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "Sub":
		err := Sub(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "GetIndexPrice":
		err := GetIndexPrice(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(ctx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(ctx)
	if ok && c != nil {
		_context = c
	}
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(_os.ToBytes()),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}
	return nil
}
