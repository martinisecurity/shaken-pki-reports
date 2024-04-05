# STIR/SHAKEN CA Ecosystem Compliance

## Sansay

Name: e_atis_ext_crl_distribution_struct\
Source: ATIS1000080\
Citation: ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements\
Effective Date: 16 Jan 22 00:00 UTC\
Description: STI intermediate and end-entity certificates shall contain a CRL Distribution Points extension containing a single DistributionPoint entry. The DistributionPoint entry shall contain a distributionPoint field identifying the HTTP URL reference to the file containing the SHAKEN CRL hosted by the STI-PA, and a cRLIssuer field that contains the DN of the issuer of the CRL.

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| error | SHAKEN 599J Technology Innovation Lab | [view](../../CERTS/1f4f7ee438c7fb0441e8b67a27ee334e35a339112bfb63acb3e8cece2322c1ac/README.md) | CRL Distribution Point shall contain a CRLIssuer field |
| error | SHAKEN 732K Serius Network | [view](../../CERTS/942253ee2461223f4489b0996e387e0a8b43f365cdbc33f791f3542c153c7dde/README.md) | CRL Distribution Point shall contain a CRLIssuer field |
| error | SHAKEN Convoso 758J | [view](../../CERTS/92c738a06590f4700891af189907fb42b4194ea6abcffc4207393c58efe3ccbe/README.md) | CRL Distribution Point shall contain a CRLIssuer field |
| error | SHAKEN 521K Voice SY LLC | [view](../../CERTS/d4f2a58cc39d7cb3d3f083c5637c392dc2972a0653aca6218501df4337a25891/README.md) | CRL Distribution Point shall contain a CRLIssuer field |

### CA Certificates

No error, warning, or notice level issues were found


Generated: 05 Apr 24 19:04 UTC