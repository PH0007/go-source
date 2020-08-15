#%%
import pandas as pd
import numpy as np

ser = pd.Series([10,2,2,3,5,4,5,9,5,6,4,5,1])
print(ser.plot.kde())