class vote:
    def __init__(self):
        self.voteaddress=set()
        self.candidate={}
        #this function is keep track of the candidates like to add the candidates
    def candidatecheck(self,candidates):
        if candidates in self.candidate:
            print("candidate already exists")
        else:
            self.candidate[candidates]=0
            #this function is to check the vote address and the candidate if the vote address is already voted or not and the candidate is exist or not 
    def votecheck(self,voteaddress,candidates):
        if voteaddress in self.voteaddress:
            print("you have already voted")
            return
        if candidates not in self.candidate:
            print("candidate does not exist")
        self.candidate[candidates]+=1
        self.voteaddress.add(voteaddress)
        #this function is to show the result
    def result(self):
        for candidate,votes in self.candidate.items():
            print(f"{candidate}: {votes} votes")
        max_vote=max(self.candidate.values())
        for candidate,votes in self.candidate.items():
            if votes == max_vote:
                print(f"Winner: {candidate}")

voting=vote()
voting.candidatecheck("tarakram")
voting.candidatecheck("mahesh")
voting.votecheck("voter1","tarakram")
voting.votecheck("voter2","mahesh")
voting.votecheck("voter1","mahesh")
voting.votecheck("voter3","tarakram" )
voting.result()