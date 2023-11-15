# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 991T

Tested At: 15 Nov 23 16:01 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 109 day(s)\
Subject: CN=SHAKEN 991T, OU=VOIP, O=kmsUScertco, L=RESTON, ST=US, C=US\
Issuer: CN=Neustar UAT Certified Caller ID SHAKEN CA-2, OU=www.ccid-uat.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cr-uat.ccid.neustar.biz/ccid/authn/v2/certs/11646.10221.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDOTCCAt6gAwIBAgIUddIRAstx9ep3WV%2Fu0K396uD0DaswCgYIKoZIzj0EAwIwgY0xCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEdMBsGA1UECwwUd3d3LmNjaWQtdWF0Lm5ldXN0YXIxNDAyBgNVBAMMK05ldXN0YXIgVUFUIENlcnRpZmllZCBDYWxsZXIgSUQgU0hBS0VOIENBLTIwHhcNMjMwMzAzMTc1MzM4WhcNMjQwMzAyMTc1MzM4WjBmMQswCQYDVQQGEwJVUzELMAkGA1UECAwCVVMxDzANBgNVBAcMBlJFU1RPTjEUMBIGA1UECgwLa21zVVNjZXJ0Y28xDTALBgNVBAsMBFZPSVAxFDASBgNVBAMMC1NIQUtFTiA5OTFUMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAED0AtcwrA%2Bc8F7%2FtOtj4mvuFTKqcbVK6seJGMrDa9VtwkAd0gYulh3A4yxbScG4hk2CkFBffooWFFGwhgn0Sb%2FKOCAUAwggE8MBYGCCsGAQUFBwEaBAowCKAGFgQ5OTFUMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUZCfMYVZl8e6kDnZKkGWHXsyi1jEwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGqBgNVHR8EgaIwgZ8wgZygPqA8hjpodHRwczovL2F1dGhlbnRpY2F0ZS1hcGktc3RnLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwHQYDVR0OBBYEFBrj51WhJLN8aR4dtnBmNFePE2iyMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEA6sw43McchTYgN8XJjC%2Bnf6adnlk8yTYVmXJi3sBJ%2FzUCIQDBtFgrKWGT26S5Nj%2Fsel0Xea1nZD%2FJcJrP5UejRmHTdQ%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 15 Nov 23 16:51 UTC