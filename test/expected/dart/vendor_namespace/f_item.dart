// Autogenerated by Frugal Compiler (2.22.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

import 'dart:typed_data' show Uint8List;
import 'package:thrift/thrift.dart' as thrift;
import 'package:vendor_namespace/vendor_namespace.dart' as t_vendor_namespace;

class Item implements thrift.TBase {
  static final thrift.TStruct _STRUCT_DESC = new thrift.TStruct("Item");



  Item() {
  }

  getFieldValue(int fieldID) {
    switch (fieldID) {
      default:
        throw new ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  setFieldValue(int fieldID, Object value) {
    switch(fieldID) {
      default:
        throw new ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  // Returns true if the field corresponding to fieldID is set (has been assigned a value) and false otherwise
  bool isSet(int fieldID) {
    switch(fieldID) {
      default:
        throw new ArgumentError("Field $fieldID doesn't exist!");
    }
  }

  read(thrift.TProtocol iprot) {
    thrift.TField field;
    iprot.readStructBegin();
    while(true) {
      field = iprot.readFieldBegin();
      if(field.type == thrift.TType.STOP) {
        break;
      }
      switch(field.id) {
        default:
          thrift.TProtocolUtil.skip(iprot, field.type);
          break;
      }
      iprot.readFieldEnd();
    }
    iprot.readStructEnd();

    // check for required fields of primitive type, which can't be checked in the validate method
    validate();
  }

  write(thrift.TProtocol oprot) {
    validate();

    oprot.writeStructBegin(_STRUCT_DESC);
    oprot.writeFieldStop();
    oprot.writeStructEnd();
  }

  String toString() {
    StringBuffer ret = new StringBuffer("Item(");

    ret.write(")");

    return ret.toString();
  }

  bool operator ==(Object o) {
    if(o == null || !(o is Item)) {
      return false;
    }
    return true;
  }

  int get hashCode {
    var value = 17;
    return value;
  }

  Item clone() {
    return new Item();
  }

  validate() {
    // check for required fields
    // check that fields of type enum have valid values
  }
}
