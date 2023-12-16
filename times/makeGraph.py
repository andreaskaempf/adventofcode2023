# Make a bar graph of times.csv

import pandas as pd
from plotly import express as px

df = pd.read_csv("times.csv")
df['Year'] = df['Year'].astype(str) # so bars side-by-side
fig = px.bar(df, x="Day", y="Minutes", color="Year", barmode="group")
fig.show()
