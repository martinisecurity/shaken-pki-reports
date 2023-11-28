# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Hawaiian Telcom SHAKEN Cert 009G

Tested At: 28 Nov 23 15:59 UTC\
Initial Validity Period: 1095 day(s)\
Remaining Validity Period: 209 day(s)\
Subject: CN=Hawaiian Telcom SHAKEN Cert 009G, O=Hawaiian Telcom, C=US\
Issuer: CN=Metaswitch STI-CA SHAKEN Issuing 1\
Link: https://cdn-cr.cgah.tnsi.com/certs/3e5c8831f88b085823dc347977abc7c3b3c4ae74

[View certificate details](https://understandingwebpki.com/?cert=MIICWDCCAf%2BgAwIBAgIQOQUqOZEeJjGXLniu6d23pTAKBggqhkjOPQQDAjAtMSswKQYDVQQDDCJNZXRhc3dpdGNoIFNUSS1DQSBTSEFLRU4gSXNzdWluZyAxMB4XDTIxMDYyNDIyMzMzN1oXDTI0MDYyMzIyMzMzN1owUjELMAkGA1UEBhMCVVMxGDAWBgNVBAoMD0hhd2FpaWFuIFRlbGNvbTEpMCcGA1UEAwwgSGF3YWlpYW4gVGVsY29tIFNIQUtFTiBDZXJ0IDAwOUcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARcJg18PBO4J1HRk8b8eGIxwzSmgW3P3K0v5EUBtwGzInfo5br1N6GzxMzBADLR49cDHXhIq9qTrVab%2FZ7v4dCZo4HbMIHYMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgXgMBYGCCsGAQUFBwEaBAowCKAGFgQwMDlHMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFPUe76ZUTXN8P8fXUWX5QnJmpNecMB8GA1UdIwQYMBaAFM0epwAQENoyHWkaOdXSRgssPIfWMAoGCCqGSM49BAMCA0cAMEQCIGI0KLpElJY3QpGIh8pEBmkDbdQp4SEtACyyHiLv1GjpAiB%2BbuEUXQfWBUZbGLFlq6j80fkNR7e%2FYzGr9pdDg3J76A%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_ext_key_usage_ee](../../ISSUES/e_atis_ext_key_usage_ee/README.md) | error | ATIS1000080 | The Key Usage extension for STI end-entity certificates shall contain a single key usage value of digitalSignature (0). |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_subject_cn_spc
- e_atis_subject_key_identifier_size
- e_atis_subject_o_required
- e_atis_tn_auth_list_spc_format
- e_shaken_certificate_policies_id

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 16:15 UTC