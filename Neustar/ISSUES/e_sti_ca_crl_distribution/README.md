# STIR/SHAKEN CA Ecosystem Compliance

## Neustar
Name: e_sti_ca_crl_distribution\
Source: ATIS-1000080v4\
Citation: ATIS-1000080.v004 / 6.4.1 STI Certificate Requirements\
Effective Date: 18 Oct 21 00:00 UTC\
Description: STI intermediate certificates shall contain a CRL Distribution Points extension containing a single DistributionPoint entry

### Leaf Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|

no warning, or error, or not effective date level issues were found

### CA Certificates

| Status | Subject | Link | Details |
|--------|---------|------|---------|
| not effective | C=US, O=Neustar Information Services Inc, OU=www.ccid.neustar, CN=Neustar Certified Caller ID CA-1 | [view](../../CERTIFICATES/9c8e8d386d404f4c97d4c37358f2362f01c9cc10/README.md) |  |
| not effective | CN=Neustar Canada Certified Caller ID SHAKEN CA-1, OU=www.ca.ccid.neustar, O=Neustar Information Services Inc, C=CA | [view](../../CERTIFICATES/a754e630241fb966de4cac2cd9eb9db90c9d029e/README.md) |  |
| not effective | CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US | [view](../../CERTIFICATES/b6f33eebd6fa1f397a22fe4d6300df28960f3061/README.md) |  |
| error | CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Inc a TransUnion company, C=US | [view](../../CERTIFICATES/aef2851bc3a7c530c6bcbf864edaeb635cb67155/README.md) | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| error | CN=Neustar Certified Caller ID SHAKEN CA-2, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US | [view](../../CERTIFICATES/3bbda8c5ef216e5cd2dc6c618a8ebd103aa90e20/README.md) | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |


Generated: 27/10/2022 at 18:57:26