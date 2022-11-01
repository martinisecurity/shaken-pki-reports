# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Neustar Certified Caller ID SHAKEN Root CA

Tested At: 01 Nov 22 16:30 UTC\
Initial Validity Period: 7305 day(s)\
Remaining Validity Period: 6865 day(s)\
Subject: CN=Neustar Certified Caller ID SHAKEN Root CA, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN Root CA, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US

View: [Click to view](https://understandingwebpki.com/?cert=MIICdzCCAh2gAwIBAgIUDHkUH5DRzTJwQ8rJF11TI%2Bl%2BxSMwCgYIKoZIzj0EAwIwgYgxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEzMDEGA1UEAwwqTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBSb290IENBMB4XDTIxMDgxNzE3MTkzN1oXDTQxMDgxNzE3MTkzN1owgYgxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEzMDEGA1UEAwwqTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBSb290IENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEmRr5XQOty4fyU%2F6oRnXlxClGX%2BeZcGIs%2B1A5eFupHNfHD3vYNQHyjM6p8msE1eUH%2FYx9Q%2BHK1W79C%2FbSHvhqwaNjMGEwDwYDVR0TAQH%2FBAUwAwEB%2FzAfBgNVHSMEGDAWgBQU1bHiD0PbARLyjgA6UwpcvXPX0DAdBgNVHQ4EFgQUFNWx4g9D2wES8o4AOlMKXL1z19AwDgYDVR0PAQH%2FBAQDAgGGMAoGCCqGSM49BAMCA0gAMEUCIQDo7VuHVpyxGw8Na0%2FjOCan1sN5I2XBuXKic5ReCMKaJwIgAey%2FdJ8o3IZlzEEdAx13ApIzzds5vdf%2FaXjz4502NSI%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ca_key_usage_crl_sign](../../ISSUES/e_us_cp_ca_key_usage_crl_sign/README.md) | error | United States SHAKEN CP | The model for managing and communicating the status of revoked certificates is in the form of a distributed Certificate Revocation List (CRL) that is maintained by the STI-PA as described in ATIS-1000080 |
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


### Not Effective

- e_atis_ca_serial_number
- e_atis_ca_subject_cn
- e_atis_root_certificate_policies
- e_atis_root_extension_unknown

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 01/11/2022 at 16:30:07