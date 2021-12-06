import createVMModel from './createVMModel';
import userModel from './userModel';
const {
  formField: {
    VM_Name,
      CPU,
      Disk,
      Memory,
      Template,
      SSHName, 
      SSHPassword
  }
} = createVMModel;

const {
  formField: {
    Email,
    Password,
  }
} = userModel;

export default {
  [VM_Name.name]: '',
  [CPU.name]: '',
  [Disk.name]: '',
  [Memory.name]: '',
  [Template.name]: '',
  [SSHName.name]: '',
  [SSHPassword.name]: '',
  [Email.name]: '',
  [Password.name]: '',
};