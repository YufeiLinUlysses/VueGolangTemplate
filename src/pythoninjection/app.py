import pymongo

myclient = pymongo.MongoClient("mongodb://localhost:27017/")
mydb = myclient["thepolyglotdeveloper"]
mycol = mydb["people"]

mylist = [
  {"firstname":"Tom","lastname":"Jerry"},
  {"firstname":"Jerry","lastname":"Tom"}
]

x = mycol.insert_many(mylist)

#print list of the _id values of the inserted documents:
print(x.inserted_ids) 