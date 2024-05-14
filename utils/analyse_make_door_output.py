str_ = '''1 - Next execution of MakeDoor Smart Contract
2024-05-14 11:55:20.578 IST 0001 INFO [chaincodeCmd] chaincodeInvokeOrQuery -> Chaincode invoke successful. result: status:200 


Output of GetAllAssets - [{"Name":"door","ID":"door","Number":19},{"Name":"steel","ID":"steel","Number":399810}]

Ended this loop execution

2 - Next execution of MakeDoor Smart Contract
Error: endorsement failure during invoke. response: status:500 message:"steel is not of the required strength. Aborting" 


Output of GetAllAssets - [{"Name":"door","ID":"door","Number":19},{"Name":"steel","ID":"steel","Number":399810}]

Ended this loop execution

3 - Next execution of MakeDoor Smart Contract
Error: endorsement failure during invoke. response: status:500 message:"steel is not of the required strength. Aborting" 


Output of GetAllAssets - [{"Name":"door","ID":"door","Number":19},{"Name":"steel","ID":"steel","Number":399810}]

Ended this loop execution

4 - Next execution of MakeDoor Smart Contract
2024-05-14 11:55:28.423 IST 0001 INFO [chaincodeCmd] chaincodeInvokeOrQuery -> Chaincode invoke successful. result: status:200 


Output of GetAllAssets - [{"Name":"door","ID":"door","Number":20},{"Name":"steel","ID":"steel","Number":399800}]

Ended this loop execution

5 - Next execution of MakeDoor Smart Contract
2024-05-14 11:55:31.032 IST 0001 INFO [chaincodeCmd] chaincodeInvokeOrQuery -> Chaincode invoke successful. result: status:200 


Output of GetAllAssets - [{"Name":"door","ID":"door","Number":21},{"Name":"steel","ID":"steel","Number":399790}]

Ended this loop execution

6 - Next execution of MakeDoor Smart Contract
2024-05-14 11:55:33.653 IST 0001 INFO [chaincodeCmd] chaincodeInvokeOrQuery -> Chaincode invoke successful. result: status:200 


Output of GetAllAssets - [{"Name":"door","ID":"door","Number":22},{"Name":"steel","ID":"steel","Number":399780}]

Ended this loop execution


'''

print(f'Number of correct executions is {str_.count("chaincodeInvokeOrQuery -> Chaincode invoke successful. result: status:200")}')

print(f'Number of die casting errors is {str_.count("die Casting is not of proper strength")}')

print(f'Number of Steel Strength errors is {str_.count("steel is not of the required strength")}')
