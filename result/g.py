import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
import os
import datetime
import re
import pdb

def avgruntime():
    dirr = os.getcwd()
    ext = '.txt'
    file_dict = {}
    txtfiles = [i for i in os.listdir(dirr) if os.path.splitext(i)[1] ==ext]
    res = {}
    for f in txtfiles:
        with open(os.path.join(dirr,f)) as file_obj:
            file_dict[f] = file_obj.read()

    for k, f in file_dict.items():
        temp = re.findall(r'real\t(.+)\n', f)
        newtemp = [int(ele[0])*60 + float(ele[2:7]) for ele in temp]
        res[k[:-4]] = sum(newtemp)/len(newtemp)

    df = pd.Dataframe()
    df.columns = ['ops','']
    pdb.set_trace()

                 
            
            

def read():
    df = pd.read_csv('res.csv')
    plt.plot( 'nodes', 'speedup', data=df[df.ops=="small"], marker='o', markerfacecolor='blue', markersize=5, color='skyblue', linewidth=2, label='small')
    plt.plot( 'nodes', 'speedup', data=df[df.ops=="medium"], marker='o', markerfacecolor='red', markersize=5, color='red', linewidth=2, label='medium')
    plt.plot( 'nodes', 'speedup', data=df[df.ops=="large"], marker='o', markerfacecolor='g', markersize=5, color='g', linewidth=2, label='large')
    plt.plot( 'nodes', 'speedup', data=df[df.ops=="test"], marker='o', markerfacecolor='black', markersize=5, color='black', linewidth=2, label='test')
    plt.legend()
    plt.title("Speedup graph with different files and different threads")
    plt.xlabel("number of threads")
    plt.ylabel("speed up")
    pdb.set_trace()
    plt.savefig('peformance.png')              
                


if __name__ == '__main__':
    read()

