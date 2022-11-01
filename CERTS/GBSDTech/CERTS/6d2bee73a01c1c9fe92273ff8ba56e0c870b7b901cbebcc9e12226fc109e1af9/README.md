# STIR/SHAKEN CA Ecosystem Compliance

## Certificate GBSDTech SHAKEN Root CA

Tested At: 01 Nov 22 20:31 UTC\
Initial Validity Period: 7300 day(s)\
Remaining Validity Period: 6755 day(s)\
Subject: CN=GBSDTech SHAKEN Root CA, O=GBSDTech, L=Ft Worth, ST=Texas, C=US\
Issuer: CN=GBSDTech SHAKEN Root CA, O=GBSDTech, L=Ft Worth, ST=Texas, C=US

View: [Click to view](https://understandingwebpki.com/?cert=MIICDTCCAbSgAwIBAgIUSRtGKKwMu2%2BiD45V7KPQ0tm4Kq0wCgYIKoZIzj0EAwIwZTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMREwDwYDVQQHDAhGdCBXb3J0aDERMA8GA1UECgwIR0JTRFRlY2gxIDAeBgNVBAMMF0dCU0RUZWNoIFNIQUtFTiBSb290IENBMB4XDTIxMDUwNTE5MDUyM1oXDTQxMDQzMDE5MDUyM1owZTELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMREwDwYDVQQHDAhGdCBXb3J0aDERMA8GA1UECgwIR0JTRFRlY2gxIDAeBgNVBAMMF0dCU0RUZWNoIFNIQUtFTiBSb290IENBMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEq8RYU3WcAwgYIrxeq%2BfdHS9BBnBYgX2LgMm%2FbLxIReFzIP%2BHwaGU1R9UqDIrydttLgr%2FdEMgZaXJ%2BVyFEewLAaNCMEAwHQYDVR0OBBYEFJcXrCQ6Lz%2FRuCzJf30oZgoSPjK6MA8GA1UdEwEB%2FwQFMAMBAf8wDgYDVR0PAQH%2FBAQDAgIEMAoGCCqGSM49BAMCA0cAMEQCIEKiPQvlaAbXWbgTA30yLXHWvYwde5bKNbg6F98X8pNVAiAKMZdFxxIVHXfKmRrmN%2FGKbPR7p29iN%2B%2B2euvX04SW8Q%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_ca_subject_rdn_unknown](../../ISSUES/w_shaken_ca_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, L, and O should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


### Not Effective

- e_atis_ca_serial_number
- e_atis_ca_subject_cn
- e_atis_root_certificate_policies
- e_atis_root_extension_unknown

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 01/11/2022 at 20:31:14