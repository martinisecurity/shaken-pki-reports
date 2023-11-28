# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 991T

Tested At: 28 Nov 23 16:02 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 96 day(s)\
Subject: CN=SHAKEN 991T, OU=VOIP, O=kmsUScertco, L=RESTON, ST=US, C=US\
Issuer: CN=Neustar UAT Certified Caller ID SHAKEN CA-2, OU=www.ccid-uat.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-uat.ccid.neustar.biz/ccid/authn/v2/certs/11646.10221.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDOTCCAt6gAwIBAgIUddIRAstx9ep3WV%2Fu0K396uD0DaswCgYIKoZIzj0EAwIwgY0xCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEdMBsGA1UECwwUd3d3LmNjaWQtdWF0Lm5ldXN0YXIxNDAyBgNVBAMMK05ldXN0YXIgVUFUIENlcnRpZmllZCBDYWxsZXIgSUQgU0hBS0VOIENBLTIwHhcNMjMwMzAzMTc1MzM4WhcNMjQwMzAyMTc1MzM4WjBmMQswCQYDVQQGEwJVUzELMAkGA1UECAwCVVMxDzANBgNVBAcMBlJFU1RPTjEUMBIGA1UECgwLa21zVVNjZXJ0Y28xDTALBgNVBAsMBFZPSVAxFDASBgNVBAMMC1NIQUtFTiA5OTFUMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAED0AtcwrA%2Bc8F7%2FtOtj4mvuFTKqcbVK6seJGMrDa9VtwkAd0gYulh3A4yxbScG4hk2CkFBffooWFFGwhgn0Sb%2FKOCAUAwggE8MBYGCCsGAQUFBwEaBAowCKAGFgQ5OTFUMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUZCfMYVZl8e6kDnZKkGWHXsyi1jEwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGqBgNVHR8EgaIwgZ8wgZygPqA8hjpodHRwczovL2F1dGhlbnRpY2F0ZS1hcGktc3RnLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwHQYDVR0OBBYEFBrj51WhJLN8aR4dtnBmNFePE2iyMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEA6sw43McchTYgN8XJjC%2Bnf6adnlk8yTYVmXJi3sBJ%2FzUCIQDBtFgrKWGT26S5Nj%2Fsel0Xea1nZD%2FJcJrP5UejRmHTdQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 16:15 UTC