'use strict'; // code generated by pbf v3.2.1

var CompressionType = exports.CompressionType = {
    "GZIP": {
        "value": 0,
        "options": {}
    },
    "SNAPPY": {
        "value": 1,
        "options": {}
    }
};

// RequestCompress ========================================

var RequestCompress = exports.RequestCompress = {};

RequestCompress.read = function (pbf, end) {
    return pbf.readFields(RequestCompress._readField, {compression_type: 0, data: null}, end);
};
RequestCompress._readField = function (tag, obj, pbf) {
    if (tag === 1) obj.compression_type = pbf.readVarint();
    else if (tag === 2) obj.data = pbf.readBytes();
};
RequestCompress.write = function (obj, pbf) {
    if (obj.compression_type) pbf.writeVarintField(1, obj.compression_type);
    if (obj.data) pbf.writeBytesField(2, obj.data);
};

// ResponseCompress ========================================

var ResponseCompress = exports.ResponseCompress = {};

ResponseCompress.read = function (pbf, end) {
    return pbf.readFields(ResponseCompress._readField, {data: null}, end);
};
ResponseCompress._readField = function (tag, obj, pbf) {
    if (tag === 1) obj.data = pbf.readBytes();
};
ResponseCompress.write = function (obj, pbf) {
    if (obj.data) pbf.writeBytesField(1, obj.data);
};

// Sdk ========================================

var Sdk = exports.Sdk = {};

Sdk.read = function (pbf, end) {
    return pbf.readFields(Sdk._readField, {code: "", source: "", name: "", version: ""}, end);
};
Sdk._readField = function (tag, obj, pbf) {
    if (tag === 1) obj.code = pbf.readString();
    else if (tag === 2) obj.source = pbf.readString();
    else if (tag === 3) obj.name = pbf.readString();
    else if (tag === 4) obj.version = pbf.readString();
};
Sdk.write = function (obj, pbf) {
    if (obj.code) pbf.writeStringField(1, obj.code);
    if (obj.source) pbf.writeStringField(2, obj.source);
    if (obj.name) pbf.writeStringField(3, obj.name);
    if (obj.version) pbf.writeStringField(4, obj.version);
};

// Os ========================================

var Os = exports.Os = {};

Os.read = function (pbf, end) {
    return pbf.readFields(Os._readField, {name: "", version: ""}, end);
};
Os._readField = function (tag, obj, pbf) {
    if (tag === 1) obj.name = pbf.readString();
    else if (tag === 2) obj.version = pbf.readString();
};
Os.write = function (obj, pbf) {
    if (obj.name) pbf.writeStringField(1, obj.name);
    if (obj.version) pbf.writeStringField(2, obj.version);
};

// Network ========================================

var Network = exports.Network = {};

Network.read = function (pbf, end) {
    return pbf.readFields(Network._readField, {cellular: false, bluetooth: false, wifi: false, address: ""}, end);
};
Network._readField = function (tag, obj, pbf) {
    if (tag === 1) obj.cellular = pbf.readBoolean();
    else if (tag === 2) obj.bluetooth = pbf.readBoolean();
    else if (tag === 3) obj.wifi = pbf.readBoolean();
    else if (tag === 4) obj.address = pbf.readString();
};
Network.write = function (obj, pbf) {
    if (obj.cellular) pbf.writeBooleanField(1, obj.cellular);
    if (obj.bluetooth) pbf.writeBooleanField(2, obj.bluetooth);
    if (obj.wifi) pbf.writeBooleanField(3, obj.wifi);
    if (obj.address) pbf.writeStringField(4, obj.address);
};

// Screen ========================================

var Screen = exports.Screen = {};

Screen.read = function (pbf, end) {
    return pbf.readFields(Screen._readField, {width: 0, height: 0}, end);
};
Screen._readField = function (tag, obj, pbf) {
    if (tag === 1) obj.width = pbf.readVarint();
    else if (tag === 2) obj.height = pbf.readVarint();
};
Screen.write = function (obj, pbf) {
    if (obj.width) pbf.writeVarintField(1, obj.width);
    if (obj.height) pbf.writeVarintField(2, obj.height);
};

// Web ========================================

var Web = exports.Web = {};

Web.read = function (pbf, end) {
    return pbf.readFields(Web._readField, {browser: "", device: "", version: "", orientation: ""}, end);
};
Web._readField = function (tag, obj, pbf) {
    if (tag === 1) obj.browser = pbf.readString();
    else if (tag === 2) obj.device = pbf.readString();
    else if (tag === 3) obj.version = pbf.readString();
    else if (tag === 4) obj.orientation = pbf.readString();
};
Web.write = function (obj, pbf) {
    if (obj.browser) pbf.writeStringField(1, obj.browser);
    if (obj.device) pbf.writeStringField(2, obj.device);
    if (obj.version) pbf.writeStringField(3, obj.version);
    if (obj.orientation) pbf.writeStringField(4, obj.orientation);
};

// Device ========================================

var Device = exports.Device = {};

Device.read = function (pbf, end) {
    return pbf.readFields(Device._readField, {d_id: "", t_id: "", u_id: "", device_id: "", profile_id: "", type: "", channel: "", os: null, network: null, screen: null, locale: "", timezone: "", web: null}, end);
};
Device._readField = function (tag, obj, pbf) {
    if (tag === 1) obj.d_id = pbf.readString();
    else if (tag === 2) obj.t_id = pbf.readString();
    else if (tag === 3) obj.u_id = pbf.readString();
    else if (tag === 4) obj.device_id = pbf.readString();
    else if (tag === 5) obj.profile_id = pbf.readString();
    else if (tag === 6) obj.type = pbf.readString();
    else if (tag === 7) obj.channel = pbf.readString();
    else if (tag === 8) obj.os = Os.read(pbf, pbf.readVarint() + pbf.pos);
    else if (tag === 9) obj.network = Network.read(pbf, pbf.readVarint() + pbf.pos);
    else if (tag === 10) obj.screen = Screen.read(pbf, pbf.readVarint() + pbf.pos);
    else if (tag === 11) obj.locale = pbf.readString();
    else if (tag === 12) obj.timezone = pbf.readString();
    else if (tag === 13) obj.web = Web.read(pbf, pbf.readVarint() + pbf.pos);
};
Device.write = function (obj, pbf) {
    if (obj.d_id) pbf.writeStringField(1, obj.d_id);
    if (obj.t_id) pbf.writeStringField(2, obj.t_id);
    if (obj.u_id) pbf.writeStringField(3, obj.u_id);
    if (obj.device_id) pbf.writeStringField(4, obj.device_id);
    if (obj.profile_id) pbf.writeStringField(5, obj.profile_id);
    if (obj.type) pbf.writeStringField(6, obj.type);
    if (obj.channel) pbf.writeStringField(7, obj.channel);
    if (obj.os) pbf.writeMessage(8, Os.write, obj.os);
    if (obj.network) pbf.writeMessage(9, Network.write, obj.network);
    if (obj.screen) pbf.writeMessage(10, Screen.write, obj.screen);
    if (obj.locale) pbf.writeStringField(11, obj.locale);
    if (obj.timezone) pbf.writeStringField(12, obj.timezone);
    if (obj.web) pbf.writeMessage(13, Web.write, obj.web);
};

// DeviceInfo ========================================

var DeviceInfo = exports.DeviceInfo = {};

DeviceInfo.read = function (pbf, end) {
    return pbf.readFields(DeviceInfo._readField, {identity: null}, end);
};
DeviceInfo._readField = function (tag, obj, pbf) {
    if (tag === 1) obj.identity = DeviceInfo.Identity.read(pbf, pbf.readVarint() + pbf.pos);
};
DeviceInfo.write = function (obj, pbf) {
    if (obj.identity) pbf.writeMessage(1, DeviceInfo.Identity.write, obj.identity);
};

// DeviceInfo.Identity ========================================

DeviceInfo.Identity = {};

DeviceInfo.Identity.read = function (pbf, end) {
    return pbf.readFields(DeviceInfo.Identity._readField, {action_time: 0, sdk: null, device: null}, end);
};
DeviceInfo.Identity._readField = function (tag, obj, pbf) {
    if (tag === 1) obj.action_time = pbf.readVarint();
    else if (tag === 2) obj.sdk = Sdk.read(pbf, pbf.readVarint() + pbf.pos);
    else if (tag === 3) obj.device = Device.read(pbf, pbf.readVarint() + pbf.pos);
};
DeviceInfo.Identity.write = function (obj, pbf) {
    if (obj.action_time) pbf.writeVarintField(1, obj.action_time);
    if (obj.sdk) pbf.writeMessage(2, Sdk.write, obj.sdk);
    if (obj.device) pbf.writeMessage(3, Device.write, obj.device);
};
