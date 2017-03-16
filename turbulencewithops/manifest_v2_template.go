package turbulencewithops

const manifestV2 = `
director_uuid: REPLACE_ME
name: REPLACE_ME

stemcells:
- alias: default
  os: ubuntu-trusty
  version: latest

releases:
- name: turbulence
  version: latest
- name: bosh-warden-cpi
  version: latest

instance_groups:
- name: api
  instances: 1
  azs: []
  jobs:
  - name: turbulence_api
    release: turbulence
  - name: warden_cpi
    release: bosh-warden-cpi
  vm_type: default
  stemcell: default
  persistent_disk_type: 1GB
  networks:
  - name: private
  properties:
    turbulence_api:
      certificate: |+
        -----BEGIN CERTIFICATE-----
        MIIEMTCCAhmgAwIBAgIRAJU/R4sb343Cs8/OZg5skPAwDQYJKoZIhvcNAQELBQAw
        GjEYMBYGA1UEAxMPdHVyYnVsZW5jZUFQSUNBMB4XDTE3MDMxNjE5MDAzM1oXDTE5
        MDMxNjE5MDAzM1owGzEZMBcGA1UEAxMQdHVyYnVsZW5jZS5sb2NhbDCCASIwDQYJ
        KoZIhvcNAQEBBQADggEPADCCAQoCggEBAL8BV14RRsMgCRMsDuJc/J0D7pMlL1Mi
        dpovUwnAhKF+RY5mVfg4DMZtMghsK274jggE8BmUYiNzDv8LAgBqMO3PC/YZa6dD
        6HUs98a0e3nE13YelYL+6riivMY1tf+MUTXZIifFdhjGMfoiXfQZfRY3CRoS1fWI
        ov7+7Xjdrb8siMfcvIyvSmBR6PORzChGn1XtHmsunBz972VXd1xg1ek2I7gWQVkA
        4t9Ov4uIimfKgx6vpzfPvkVuhuJ/9qN1avTBvB00uD3jWyRmcL/+9dU9MAO+OBGE
        oxPFk75h++oB2BO9W650iiu295OgWvHM8f2/vMmyABQ0cwudet+7LWcCAwEAAaNx
        MG8wDgYDVR0PAQH/BAQDAgO4MB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcD
        AjAdBgNVHQ4EFgQUvLTn+EUJrIx9EjyX3CsaT4fJwXkwHwYDVR0jBBgwFoAUc10U
        M7akDSV7LqWGSEAxUvGK3/IwDQYJKoZIhvcNAQELBQADggIBAEwS0j3tmGbwj28a
        sTo7VQfsxG6HVVIA1TfIdYwkS+f6vFayqihuU3/cWBWC6hpbbr/Du6miHjYuwILY
        LhcPNGRuX8Zlx36GmsNmcrdoDgnZxXu0HFeMyAFEW8zf2pfjG5h989mFb5W0TqcW
        pEdCROLOgbsYGU9sKDS16XK0RoshQdfpaepIpDiteK4vH7B6uBMy4/3NiYn8JC1f
        dJyzmY8lpi3mUYJ87FjcK5WiF01gBFhbcBghEsXH6kvHQN2DtW4fvHTtzEBmJ6oa
        aeEfchi/UAk2oET2H/L2E8leYQxMFRpr2ybXVXDUp77Ur3c8N+EFPePtwk+vzX7o
        DF//iwM03hKIZkInXmeI+fTOEVBuCLxz+C8jFb59rFd5cqJdAqfpHHmkJmHaRVnf
        qx0MdkHjRPJXIgMjZVC+XbBf/gCLgeV0W5orZvyDmsGePrLKjFiVnE0vXb4h8Ggy
        2eQ8OB5lz1y9XMOYkEur2/bKCmiVL7yInsskEWBSB+yQPXpMBYOBEwIUzQFnMGQy
        cq1UBiNmr18evA5ee9M8Axcjptogz4U+9NcnQZlJpFbf/v88KAc2MoGrq8rvAHhZ
        +I+b3oKvzhPiM3WO+v+Lu7M3s7k5VZ8joIg6QnlIR+38f75tKQduoifQoyCjjnDS
        iQefcookOmhMBxZj+1YoSSZc5nSd
        -----END CERTIFICATE-----
      cpi_job_name: warden_cpi
      director:
        ca_cert: |+
          -----BEGIN CERTIFICATE-----
          MIIDtzCCAp+gAwIBAgIJAMZ/qRdRamluMA0GCSqGSIb3DQEBBQUAMEUxCzAJBgNV
          BAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBX
          aWRnaXRzIFB0eSBMdGQwIBcNMTYwODI2MjIzMzE5WhgPMjI5MDA2MTAyMjMzMTla
          MEUxCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJ
          bnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw
          ggEKAoIBAQDN/bv70wDn6APMqiJZV7ESZhUyGu8OzuaeEfb+64SNvQIIME0s9+i7
          D9gKAZjtoC2Tr9bJBqsKdVhREd/X6ePTaopxL8shC9GxXmTqJ1+vKT6UxN4kHr3U
          +Y+LK2SGYUAvE44nv7sBbiLxDl580P00ouYTf6RJgW6gOuKpIGcvsTGA4+u0UTc+
          y4pj6sT0+e3xj//Y4wbLdeJ6cfcNTU63jiHpKc9Rgo4Tcy97WeEryXWz93rtRh8d
          pvQKHVDU/26EkNsPSsn9AHNgaa+iOA2glZ2EzZ8xoaMPrHgQhcxoi8maFzfM2dX2
          XB1BOswa/46yqfzc4xAwaW0MLZLg3NffAgMBAAGjgacwgaQwHQYDVR0OBBYEFNRJ
          PYFebixALIR2Ee+yFoSqurxqMHUGA1UdIwRuMGyAFNRJPYFebixALIR2Ee+yFoSq
          urxqoUmkRzBFMQswCQYDVQQGEwJBVTETMBEGA1UECBMKU29tZS1TdGF0ZTEhMB8G
          A1UEChMYSW50ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkggkAxn+pF1FqaW4wDAYDVR0T
          BAUwAwEB/zANBgkqhkiG9w0BAQUFAAOCAQEAoPTwU2rm0ca5b8xMni3vpjYmB9NW
          oSpGcWENbvu/p7NpiPAe143c5EPCuEHue/AbHWWxBzNAZvhVZBeFirYNB3HYnCla
          jP4WI3o2Q0MpGy3kMYigEYG76WeZAM5ovl0qDP6fKuikZofeiygb8lPs7Hv4/88x
          pSsZYBm7UPTS3Pl044oZfRJdqTpyHVPDqwiYD5KQcI0yHUE9v5KC0CnqOrU/83PE
          b0lpHA8bE9gQTQjmIa8MIpaP3UNTxvmKfEQnk5UAZ5xY2at5mmyj3t8woGdzoL98
          yDd2GtrGsguQXM2op+4LqEdHef57g7vwolZejJqN776Xu/lZtCTp01+HTA==
          -----END CERTIFICATE-----
        host: REPLACE_ME
        password: REPLACE_ME
        username: REPLACE_ME
      password: turbulence-password
      private_key: |+
        -----BEGIN RSA PRIVATE KEY-----
        MIIEogIBAAKCAQEAvwFXXhFGwyAJEywO4lz8nQPukyUvUyJ2mi9TCcCEoX5FjmZV
        +DgMxm0yCGwrbviOCATwGZRiI3MO/wsCAGow7c8L9hlrp0PodSz3xrR7ecTXdh6V
        gv7quKK8xjW1/4xRNdkiJ8V2GMYx+iJd9Bl9FjcJGhLV9Yii/v7teN2tvyyIx9y8
        jK9KYFHo85HMKEafVe0eay6cHP3vZVd3XGDV6TYjuBZBWQDi306/i4iKZ8qDHq+n
        N8++RW6G4n/2o3Vq9MG8HTS4PeNbJGZwv/711T0wA744EYSjE8WTvmH76gHYE71b
        rnSKK7b3k6Ba8czx/b+8ybIAFDRzC51637stZwIDAQABAoIBADXHFuYxTw8ZMfTn
        7rjlHWrH9KARVCXACkyXDnYscitV9heF0Ka2gUJM9E1Sx1fTO4oeect577ezaYF9
        g+7B90y9gsyjk5/gis6S/qA/qJQ8S73CFq5vP38EssnLzZZJ14OlwuwXuIx5IREg
        I/vaQVHD5RgmPX+IHPxWol3pbEXqsrPcMtGGjCyuR/P379nmk8snpA/2AwgQRm/D
        p/UXRoBwhJzx4R0X8JhZWZzLapct8pxnxW0nMbA68FCesK8Wr+AfaeCPAbtWsQaT
        OFHiJk8zd3IA3MXBSCmxAgBDXlQTPGWQqFsGVhKYyvOkUzcjNSKwXG8e28WMJ0eJ
        Xq1MEUECgYEA99WTzzOnPWAVYy4DwbcF64K5RIkDNXVgFgjvTeNZg0pTWTijljYd
        BTThapHJeBQhR2B0JhYWMdE1df7cyvvk52LptccyWN4PbotuLfZkS7E7Ygisr/EA
        qStkAepTaEx5knYw3pXuOgcMS9vGoAIuRIhaFM3KvB/YCvq/GNdBRt8CgYEAxUxs
        rzl9u5mXakUyFmdleOug2tUZbzpM5vTvPdRGCcJcbBR5TKRYmsoy7HK2QaDY/XHW
        ejzMREM2Omflhl1gP1Oy7RYCyOTYIwVZSDXzyZZ4I7NdPo8zQmYeCTyXVjCSwQak
        b51mDhdApzeOKVLeDKmKvT7zmczZbRCVCvq7EnkCgYBPZQF0GIUUGWrgmgYkEcD4
        wKkfdpErmA2PIY+gMRwk/jOTWpy1a2KCn30zSb70E+bRWen7pYm4rd/ljB8pe+bJ
        5ZsfW9AaPhFNhadnXA6nXQC8GDFSL+/ZQghIwMu1lwI/VFO1iuyFdGqRBrgr3Gb8
        F35cOc0f1Ue9xLRhfdvwmwKBgBuWkCwAUgCFfZKykkpnstyxthBjc+cFs/MnlLyE
        jjaXIu1J/5wj7u/WDkDZ2Xpbz3vBC0iUb8uryk1occPUEr1IKuDUDxegSEi9Wrqq
        MKijjbEQR52T0IscVF7eRhsbN6oeD6g7ziVyQuwe7JYCrGIA4xGLV+zNpCmIBA9A
        +B+hAoGAbZ5aT+RBf3TIgrWG2pMHcGT9XkIGzyXazWyvYVnX2deW+wE98/NSoBep
        WiOuGPj6duy7Tn75XhAiXtcpLsSCqy6dgTkaG9+47jfVf4Kz32yh8e3A8LsUnekA
        35JD5rJPjgZfec6c03uP62aUqUfa90HANh8EI98OBuWc6C4MeZ4=
        -----END RSA PRIVATE KEY-----
    warden_cpi:
      agent:
        blobstore:
          options:
            endpoint: http://10.254.50.4:25251
            password: agent-password
            user: agent
          provider: dav
        mbus: nats://nats:nats-password@10.254.50.4:4222
      warden:
        connect_address: 10.254.50.4:7777
        connect_network: tcp

update:
  canaries: 1
  canary_watch_time: 1000-180000
  max_in_flight: 1
  serial: true
  update_watch_time: 1000-180000
`