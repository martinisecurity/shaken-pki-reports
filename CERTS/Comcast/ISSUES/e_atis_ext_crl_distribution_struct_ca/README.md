# STIR/SHAKEN CA Ecosystem Compliance

## Comcast

Name: e_atis_ext_crl_distribution_struct_ca\
Source: ATIS1000080\
Citation: ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements\
Effective Date: 18 Oct 21 00:00 UTC\
Description: STI intermediate and end-entity certificates shall contain a CRL Distribution Points extension containing a single DistributionPoint entry. The DistributionPoint entry shall contain a distributionPoint field identifying the HTTP URL reference to the file containing the SHAKEN CRL hosted by the STI-PA, and a cRLIssuer field that contains the DN of the issuer of the CRL.

### Leaf Certificates

No error, warning, or notice level issues were found

### CA Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | Comcast SHAKEN Intermediate CA | [view](../../CERTS/25af737667ed8b05cb8b8e7f44b2d7b5861551bea95ec48a73306ec75a92a662/README.md) | CRL Distribution Point shall contain a CRLIssuer field |


Generated: 22 Aug 24 15:44 UTC