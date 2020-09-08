FROM scratch
ADD pingpong /
EXPOSE 8080
CMD ["/pingpong"]
