const {DeviceInfo, Device, Screen, Sdk, Network, Os, Web} = require('./user_tracking_pb.js');
const {DeviceServiceClient} = require('./user_tracking_grpc_web_pb.js');
var protobuf = require("protobufjs");

module.exports = {
  run: function (data) {
    let service = new DeviceServiceClient('http://localhost:8080');
    let request = new DeviceInfo()

    let deviceInfoString = String.fromCharCode(...data);
    let deviceInfo = JSON.parse(deviceInfoString);
    console.log(deviceInfo)
    
    request.setIdentity(new DeviceInfo.Identity()
                        .setActionTime(deviceInfo['identity']['action_time'])
                        .setSdk(new Sdk()
                          .setCode(deviceInfo['identity']['sdk']['code'])
                          .setName(deviceInfo['identity']['sdk']['name'])
                          .setSource(deviceInfo['identity']['sdk']['source'])
                          .setVersion(deviceInfo['identity']['sdk']['version'])
                        )
                        .setDevice(new Device()
                          .setChannel(deviceInfo['identity']['device']['channel'])
                          .setDId(deviceInfo['identity']['device']['d_id'])
                          .setLocale(deviceInfo['identity']['device']['locale'])
                          .setDeviceId(deviceInfo['identity']['device']['device_id'])
                          .setNetwork(new Network()
                            .setAddress(deviceInfo['identity']['device']['network']['address'])
                            .setBluetooth(deviceInfo['identity']['device']['network']['bluetooth'])
                            .setCellular(deviceInfo['identity']['device']['network']['cellular'])
                            .setWifi(deviceInfo['identity']['device']['network']['wifi'])
                          )
                          .setOs(new Os()
                            .setName(deviceInfo['identity']['device']['os']['name'])
                            .setVersion(deviceInfo['identity']['device']['os']['version'])
                          )
                          .setProfileId(deviceInfo['identity']['device']['profile_id'])
                          .setScreen(new Screen()
                            .setWidth(deviceInfo['identity']['device']['screen']['width'])
                            .setHeight(deviceInfo['identity']['device']['screen']['height'])
                          )
                          .setTId(deviceInfo['identity']['device']['t_id'])
                          .setTimezone(deviceInfo['identity']['device']['timezone'])
                          .setType(deviceInfo['identity']['device']['type'])
                          .setUId(deviceInfo['identity']['device']['u_id'])
                          .setWeb(new Web()
                            .setBrowser(deviceInfo['identity']['device']['web']['browser'])
                            .setDevice(deviceInfo['identity']['device']['web']['device'])
                            .setVersion(deviceInfo['identity']['device']['web']['version'])
                            .setOrientation(deviceInfo['identity']['device']['web']['orientation'])
                          )
                        )
                      )
    console.log(request.toObject())
    // request.fromObject(deviceInfo)

    service.registerDevice(request, {}, function(err, response) {
        if (err) {
            console.log(err.code);
            console.log(err.message);
          } else {
            console.log('Response')
            console.log(response);
            console.log('Device id', response.getDId())
          }
    });
  }
};