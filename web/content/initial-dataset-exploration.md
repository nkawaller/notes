
# Initial Dataset Exploration and Analysis

  The following steps will give you a comprehensive understanding of the
  dataset's structure, content, and potential issues befor diving into
  deeper analysis or model-building.


## Load the Data Efficiently

   Load only a portion of the dataset first, especially if it's large, 
   to avoid overwhelming memory:

   ```python
   import pandas as pd

   # Load a sample of 100 rows
   df = pd.read_csv('your_file.csv', nrows=100)
   ```
   You can also load in chunks if the file is very large:
   ```python
   chunk_size = 10000
   for chunk in pd.read_csv('your_file.csv', chunksize=chunk_size):
       # Process each chunk here
       pass
   ```

## Get an Overview of the Data
   
   Use methods to quickly assess the structure:
   
   **First few rows**:
   ```python
   df.head()  # View first 5 rows
   df.tail()  # View last 5 rows
   ```
   **Summary of data types and non-null counts**:
   ```python
   df.info()
   ```
   **Column names**:
   ```python
   df.columns
   ```
   **Number of rows and columns**:
   ```python
   df.shape
   ```

## Check for Missing Data

   Missing data can affect analysis and model performance:
   ```python
   # Count missing values per column
   df.isnull().sum()
   
   # Percentage of missing values per column
   df.isnull().mean()  
   ```

## Look at Descriptive Statistics
   
   Get an overall sense of the data’s distribution:
   ```python
   # Summary statistics for numeric columns
   df.describe()
   
   # Summary for all columns, including categorical
   df.describe(include='all')
   ```

## Examine Data Types
  
  Ensure the data types (e.g., numeric, categorical) are correct:
   ```python
   df.dtypes
   ```
   If you notice inconsistencies (e.g., a column with numbers stored as 
   strings), you might want to convert them:
   ```python
   # Convert to numeric, with errors set to NaN
   df['column'] = pd.to_numeric(df['column'], errors='coerce')
   ```

## Handle Duplicates

   Check for any duplicate rows that may need to be removed:
   ```python
   # Number of duplicate rows
   df.duplicated().sum()
   
   # Remove duplicates
   df.drop_duplicates(inplace=True)
   ```

## Examine Unique Values in Categorical Columns

   For categorical data, check the number of unique values and their 
   distribution:
   ```python
   # Distribution of values
   df['column'].value_counts()
   
   # Number of unique values
   df['column'].nunique()
   ```

## Check for Outliers

   Use basic visualizations to check for outliers in numeric columns:
   ```python
   import matplotlib.pyplot as plt
   df.boxplot(column='numeric_column')
   plt.show()
   ```

## Inspect Data Relationships

   Check correlations between numerical columns to identify potential 
   relationships:
   ```python
   df.corr()
   ```

## Visualize the Data

   Simple plots can give you insights into patterns:
   
   **Histograms for distributions**:
   ```python
   df['numeric_column'].hist()
   plt.show()
   ```
   **Bar charts for categorical data**:
   ```python
   df['categorical_column'].value_counts().plot(kind='bar')
   plt.show()
   ```

## Memory Usage

   For large datasets, monitor memory usage to avoid crashes:
   ```python
   # Total memory usage in bytes
   df.memory_usage(deep=True).sum()
   ```
   Consider optimizing memory by converting data types (e.g., converting 
   large integer columns to smaller types):
   ```python
   # Convert to a smaller int type
   df['column'] = df['column'].astype('int32')
   ```
