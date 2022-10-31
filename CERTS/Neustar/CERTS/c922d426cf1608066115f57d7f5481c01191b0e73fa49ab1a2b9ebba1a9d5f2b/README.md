# STIR/SHAKEN CA Ecosystem Compliance

## Certificate prod SHAKEN 811J

Tested At: 31 Oct 22 16:39 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 25 day(s)\
Subject: CN=prod SHAKEN 811J, O=Alianza, L=Pleasant Grove, ST=Utah, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://api.alianza.com/v2/stir-shaken/certs/b45a4083-1554-4412-b5fc-bbd2c027091e/key.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDJTCCAsugAwIBAgIUF11AhAdij74OllsQYOfJXA13UAMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIxMTEyNDE5MDgyM1oXDTIyMTEyNDE5MDgyM1owYjELMAkGA1UEBhMCVVMxDTALBgNVBAgMBFV0YWgxFzAVBgNVBAcMDlBsZWFzYW50IEdyb3ZlMRAwDgYDVQQKDAdBbGlhbnphMRkwFwYDVQQDDBBwcm9kIFNIQUtFTiA4MTFKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAElnn93NPvWDVuPqwD8TmCh9ozF4j6OCbTcFrObVZ6DZ%2FS%2FWcodSyA%2BkugUN1PYWrP2M%2FwzgFajAgJYdNGMcrBW6OCATkwggE1MBYGCCsGAQUFBwEaBAowCKAGFgQ4MTFKMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUr9HIwu5yTIP8P%2B0Zp20dkLIH8DowWwYIKwYBBQUHAQEETzBNMEsGCCsGAQUFBzAChj9odHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydCAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAdBgNVHQ4EFgQU5zlnG7MzIXeihn%2FxrA8I8PhsiLYwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIDcsPV9TNLAyel95AnHjSHqp0vs%2FONtwHftmY%2BeuoYuDAiEA%2Ffp9UnRdPfq5qPRz249s6r3yzJAWkX7qlBrxT90UKuQ%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


### Not Effective

- e_cp1_3_ambiguous_identifier
- e_cp1_3_subject_sn
- e_sti_extension_unknown
- e_sti_serial_number
- e_sti_signature_algorithm
- e_sti_subject_cn
- w_cp_1_3_subject_email

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 31/10/2022 at 16:43:22