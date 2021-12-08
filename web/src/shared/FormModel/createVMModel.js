export default {
    formId: 'createVMModel',
    formField: {
        VM_Name: {
        name: 'VM_Name',
        label: 'Virtual Machine Name',
        requiredErrorMsg: 'Is required'
      },
      CPU: {
        name: 'CPUs',
        label: 'CPU*',
        requiredErrorMsg: 'Is required'
      },
      Disk: {
        name: 'Disk',
        label: 'Hard Disk*',
        requiredErrorMsg: 'Is required'
      },
      Memory: {
        name: 'Memory',
        label: 'RAM'
      },
      Template: {
        name: 'Template',
        label: 'OS Template*',
        data: [{
          value: 'ubuntu16.4',
          label: 'Ubuntu 16.4'
        },
        {
          value: 'centos',
          label: 'CentOS 3'
        }],
        requiredErrorMsg: 'Is required'
      },
      SSHName: {
        name: 'SSH_username',
        label: 'SSH Name*',
        requiredErrorMsg: 'Is required',
      },
      SSHPassword: {
        name: 'SSH_password',
        label: 'SSH Password*',
        requiredErrorMsg: 'Is required',
      },
    }
  };
  