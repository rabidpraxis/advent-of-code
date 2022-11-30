(ns advent.utils
  (:require
    [clojure.string :as s]
    [clojure.java.io :as io]))

(defn get-input [name]
  (slurp (io/resource (str name ".txt"))))

(defn get-input-lines [name]
  (s/split (get-input name) #"\n"))
